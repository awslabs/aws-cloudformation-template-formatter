package format

import (
	"fmt"
	"sort"
	"strings"
)

func formatMap(data map[string]interface{}) string {
	if len(data) == 0 {
		return "{}"
	}

	keys := make([]string, 0, len(data))

	for key, _ := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	parts := make([]string, len(keys))

	for i, key := range keys {
		value := data[key]
		fmtValue := Yaml(value)

		switch value.(type) {
		case map[string]interface{}:
			parts[i] = fmt.Sprintf("%s:\n  %s", key, strings.Replace(fmtValue, "\n", "\n  ", -1))
		default:
			parts[i] = fmt.Sprintf("%s: %s", key, fmtValue)
		}
	}

	return strings.Join(parts, "\n")
}

func formatList(data []interface{}) string {
	if len(data) == 0 {
		return "[]"
	}

	parts := make([]string, len(data))

	for i, value := range data {
		parts[i] = fmt.Sprintf("- %s", strings.Replace(Yaml(value), "\n", "\n  ", -1))
	}

	return strings.Join(parts, "\n")
}

func Yaml(data interface{}) string {
	switch value := data.(type) {
	case map[string]interface{}:
		return formatMap(value)
	case []interface{}:
		return formatList(value)
	default:
		return fmt.Sprint(value)
	}
}
