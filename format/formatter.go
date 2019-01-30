package format

import (
	"fmt"
	"strings"
)

type formatter struct {
	style          string
	data           value
	path           []interface{}
	currentValue   interface{}
	currentComment string
}

func newFormatterWithComments(style string, data interface{}, comments map[interface{}]interface{}) formatter {
	p := formatter{
		style: style,
		data: value{
			data,
			comments,
		},
		path: make([]interface{}, 0),
	}

	p.get()

	return p
}

func newFormatter(style string, data map[string]interface{}) formatter {
	return newFormatterWithComments(style, data, nil)
}

func (p *formatter) get() {
	p.currentValue = p.data.Get(p.path)
	p.currentComment = p.data.GetComment(p.path)
}

func (p *formatter) push(key interface{}) {
	p.path = append(p.path, key)
	p.get()
}

func (p *formatter) pop() {
	p.path = p.path[:len(p.path)-1]
	p.get()
}

func (p formatter) indent(in string) string {
	indenter := "  "

	if p.style == "json" {
		indenter = "    "
	}
	parts := strings.Split(in, "\n")

	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = indenter + part
		}
	}

	if p.style == "json" {
		return strings.Join(parts, "\n")
	}

	return strings.TrimLeft(strings.Join(parts, "\n"), " ")
}

func (p formatter) formatIntrinsic(key string) string {
	p.push(key)
	defer p.pop()

	if p.style == "json" {
		return p.format()
	}

	shortKey := strings.Replace(key, "Fn::", "", 1)

	fmtValue := p.format()

	switch p.currentValue.(type) {
	case map[string]interface{}, []interface{}:
		return fmt.Sprintf("!%s\n  %s", shortKey, p.indent(fmtValue))
	default:
		return fmt.Sprintf("!%s %s", shortKey, fmtValue)
	}
}

func (p formatter) formatMap(data map[string]interface{}) string {
	if len(data) == 0 {
		return "{}"
	}

	keys := sortKeys(data, p.path)

	parts := make([]string, len(keys))

	for i, key := range keys {
		value := data[key]

		p.push(key)
		fmtValue := p.format()

		if p.style == "json" {
			fmtValue = fmt.Sprintf("%q: %s", key, fmtValue)
			if i < len(keys)-1 {
				fmtValue += ","
			}

			isScalar := true
			switch value.(type) {
			case map[string]interface{}, []interface{}:
				isScalar = false
			}

			if p.currentComment != "" && isScalar {
				fmtValue += "  // " + p.currentComment
			}
		} else {
			needsIndent := false

			switch v := value.(type) {
			case map[string]interface{}:
				if iKey, ok := intrinsicKey(v); ok {
					fmtValue = p.formatIntrinsic(iKey)
				} else {
					needsIndent = true
				}
			case []interface{}:
				if fmtValue != "[]" {
					needsIndent = true
				}
			}

			if needsIndent {
				if p.currentComment != "" {
					fmtValue = fmt.Sprintf("%s:  # %s\n  %s", key, p.currentComment, p.indent(fmtValue))
				} else {
					fmtValue = fmt.Sprintf("%s:\n  %s", key, p.indent(fmtValue))
				}
			} else {
				if p.currentComment != "" {
					fmtValue = fmt.Sprintf("%s: %s  # %s", key, fmtValue, p.currentComment)
				} else {
					fmtValue = fmt.Sprintf("%s: %s", key, fmtValue)
				}
			}
		}

		parts[i] = fmtValue

		p.pop()
	}

	// Double gap for top-level elements
	joiner := "\n"
	if len(p.path) <= 1 {
		joiner = "\n\n"
	}

	if p.style == "json" {
		if p.currentComment != "" {
			return "{  // " + p.currentComment + "\n" + p.indent(strings.Join(parts, joiner)) + "\n}"
		}

		return "{\n" + p.indent(strings.Join(parts, joiner)) + "\n}"
	}

	return strings.Join(parts, joiner)
}

func (p formatter) formatList(data []interface{}) string {
	if len(data) == 0 {
		return "[]"
	}

	parts := make([]string, len(data))

	for i := range data {
		p.push(i)
		fmtValue := p.format()

		if p.style == "json" {
			parts[i] = p.indent(fmtValue)
		} else {
			parts[i] = fmt.Sprintf("- %s", p.indent(fmtValue))
		}

		if p.currentComment != "" {
			if p.style == "json" {
				parts[i] += "  // " + p.currentComment
			} else {
				parts[i] += "  # " + p.currentComment
			}
		}

		p.pop()
	}

	if p.style == "json" {
		if p.currentComment != "" {
			return "[  // " + p.currentComment + "\n" + strings.Join(parts, ",\n") + "\n]"
		}

		return "[\n" + strings.Join(parts, ",\n") + "\n]"
	}

	return strings.Join(parts, "\n")
}

func (p formatter) format() string {
	switch v := p.currentValue.(type) {
	case map[string]interface{}:
		return p.formatMap(v)
	case []interface{}:
		return p.formatList(v)
	case string:
		if p.style == "json" {
			return fmt.Sprintf("%q", v)
		}

		return formatString(v)
	default:
		return fmt.Sprint(v)
	}
}
