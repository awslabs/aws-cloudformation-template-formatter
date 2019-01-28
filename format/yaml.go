package format

import (
	"fmt"
	"strings"
)

type formatter struct {
	style          string
	data           map[string]interface{}
	comments       map[string]interface{}
	path           []interface{}
	currentValue   interface{}
	currentComment interface{}
}

func newFormatterWithComments(style string, data, comments map[string]interface{}) formatter {
	return formatter{
		style,
		data,
		comments,
		make([]interface{}, 0),
		data,
		comments,
	}
}

func newFormatter(style string, data map[string]interface{}) formatter {
	return newFormatterWithComments(style, data, nil)
}

func (p *formatter) push(key interface{}) {
	p.path = append(p.path, key)
	p.currentValue = mustGetFromPath(p.data, p.path)
	p.currentComment = getFromPath(p.comments, p.path)
}

func (p *formatter) pop() {
	p.path = p.path[:len(p.path)-1]
	p.currentValue = mustGetFromPath(p.data, p.path)
	p.currentComment = getFromPath(p.comments, p.path)
}

func (p formatter) formatIntrinsic(key string) string {
	p.push(key)
	defer p.pop()

	shortKey := strings.Replace(key, "Fn::", "", 1)

	fmtValue := p.format()

	switch p.currentValue.(type) {
	case map[string]interface{}, []interface{}:
		return fmt.Sprintf("!%s\n  %s", shortKey, indent(fmtValue))
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
			if comment, ok := p.currentComment.(string); ok {
				fmtValue = fmt.Sprintf("%s:  # %s\n  %s", key, comment, indent(fmtValue))
			} else {
				fmtValue = fmt.Sprintf("%s:\n  %s", key, indent(fmtValue))
			}
		} else {
			if comment, ok := p.currentComment.(string); ok {
				fmtValue = fmt.Sprintf("%s: %s  # %s", key, fmtValue, comment)
			} else {
				fmtValue = fmt.Sprintf("%s: %s", key, fmtValue)
			}
		}

		parts[i] = fmtValue

		p.pop()
	}

	joiner := "\n"

	if len(p.path) <= 1 {
		joiner = "\n\n"
	}

	return strings.Join(parts, joiner)
}

func (p formatter) formatList(data []interface{}) string {
	if len(data) == 0 {
		return "[]"
	}

	parts := make([]string, len(data))

	for i, _ := range data {
		p.push(i)
		fmtValue := p.format()
		p.pop()

		parts[i] = fmt.Sprintf("- %s", indent(fmtValue))
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
		return formatString(v)
	default:
		return fmt.Sprint(v)
	}
}

func Yaml(data map[string]interface{}) string {
	return newFormatter("yaml", data).format()
}

func YamlWithComments(data, comments map[string]interface{}) string {
	return newFormatterWithComments("yaml", data, comments).format()
}
