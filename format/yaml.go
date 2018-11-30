package format

import (
	"fmt"
	"sort"
	"strings"
)

func indent(in string) string {
	parts := strings.Split(in, "\n")

	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = "  " + part
		}
	}

	return strings.TrimLeft(strings.Join(parts, "\n"), " ")
}

func sortKeys(data map[string]interface{}, path []string) []string {
	// See if we have some required-order keys
	order := ordering(path)

	done := make(map[string]bool)
	orderedKeys := make([]string, 0, len(data))
	otherKeys := make([]string, 0, len(data))

	// Apply required keys
	for _, orderedKey := range order {
		if _, ok := data[orderedKey]; ok {
			orderedKeys = append(orderedKeys, orderedKey)
			done[orderedKey] = true
		}
	}

	// Now the remainder of the keys
	for key, _ := range data {
		if !done[key] {
			otherKeys = append(otherKeys, key)
		}
	}
	sort.Strings(otherKeys)

	return append(orderedKeys, otherKeys...)
}

func formatMap(data map[string]interface{}, path []string) string {
	if len(data) == 0 {
		return "{}"
	}

	keys := sortKeys(data, path)

	parts := make([]string, len(keys))

	for i, key := range keys {
		value := data[key]
		fmtValue := yaml(value, append(path, key))

		switch value.(type) {
		case map[string]interface{}:
			fmtValue = fmt.Sprintf("%s:\n  %s", key, indent(fmtValue))
		case []interface{}:
			fmtValue = fmt.Sprintf("%s:\n  %s", key, indent(fmtValue))
		default:
			fmtValue = fmt.Sprintf("%s: %s", key, fmtValue)
		}

		parts[i] = fmtValue
	}

	output := strings.Join(parts, "\n")

	if len(path) == 2 {
		output += "\n"
	}

	return output
}

func formatList(data []interface{}, path []string) string {
	if len(data) == 0 {
		return "[]"
	}

	parts := make([]string, len(data))

	for i, value := range data {
		fmtValue := yaml(value, append(path, string(i)))

		parts[i] = fmt.Sprintf("- %s", indent(fmtValue))
	}

	return strings.Join(parts, "\n")
}

func yaml(data interface{}, path []string) string {
	switch value := data.(type) {
	case map[string]interface{}:
		return formatMap(value, path)
	case []interface{}:
		return formatList(value, path)
	default:
		return fmt.Sprint(value)
	}
}

func Yaml(data interface{}) string {
	return yaml(data, make([]string, 0))
}
