package format

import (
	"fmt"
	"strings"
)

func formatJsonString(data string) string {
	return fmt.Sprintf("%q", data)
}

func indentJson(in string) string {
	parts := strings.Split(in, "\n")

	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = "    " + part
		}
	}

	return strings.Join(parts, "\n")
}

func formatJsonIntrinsic(key string, data interface{}, path []interface{}) string {
	shortKey := strings.Replace(key, "Fn::", "", 1)

	fmtValue := json(data, path)

	switch data.(type) {
	case map[string]interface{}:
		return fmt.Sprintf("!%s\n  %s", shortKey, indentJson(fmtValue))
	case []interface{}:
		return fmt.Sprintf("!%s\n  %s", shortKey, indentJson(fmtValue))
	default:
		return fmt.Sprintf("!%s %s", shortKey, json(data, path))
	}
}

func formatJsonMap(data map[string]interface{}, path []interface{}) string {
	if len(data) == 0 {
		return "{}"
	}

	keys := sortKeys(data, path)

	parts := make([]string, len(keys))

	for i, key := range keys {
		value := data[key]
		fmtValue := json(value, append(path, key))
		fmtValue = fmt.Sprintf("%s: %s", formatJsonString(key), fmtValue)
		parts[i] = fmtValue
	}

	joiner := ",\n"

	return "{\n" + indentJson(strings.Join(parts, joiner)) + "\n}"
}

func formatJsonList(data []interface{}, path []interface{}) string {
	if len(data) == 0 {
		return "[]"
	}

	parts := make([]string, len(data))

	for i, value := range data {
		fmtValue := json(value, append(path, i))

		parts[i] = indentJson(fmtValue)
	}

	return "[\n" + strings.Join(parts, ",\n") + "\n]"
}

func json(data interface{}, path []interface{}) string {
	if value, ok := data.(map[string]interface{}); ok {
		return formatJsonMap(value, path)
	}

	if value, ok := data.([]interface{}); ok {
		return formatJsonList(value, path)
	}

	if value, ok := data.(string); ok {
		return formatJsonString(value)
	}

	return fmt.Sprint(data)
}

func Json(data interface{}) string {
	return json(data, make([]interface{}, 0))
}
