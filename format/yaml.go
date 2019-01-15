package format

import (
	"fmt"
	"strings"
)

func formatYamlIntrinsic(key string, data interface{}, path []string) string {
	shortKey := strings.Replace(key, "Fn::", "", 1)

	fmtValue := yaml(data, path)

	switch data.(type) {
	case map[string]interface{}:
		return fmt.Sprintf("!%s\n  %s", shortKey, indent(fmtValue))
	case []interface{}:
		return fmt.Sprintf("!%s\n  %s", shortKey, indent(fmtValue))
	default:
		return fmt.Sprintf("!%s %s", shortKey, yaml(data, path))
	}
}

func formatYamlMap(data map[string]interface{}, path []string) string {
	if len(data) == 0 {
		return "{}"
	}

	keys := sortKeys(data, path)

	parts := make([]string, len(keys))

	for i, key := range keys {
		value := data[key]
		fmtValue := yaml(value, append(path, key))

		switch v := value.(type) {
		case map[string]interface{}:
			if iKey, ok := intrinsicKey(v); ok {
				fmtValue = formatYamlIntrinsic(iKey, v[iKey], append(path, key))
				fmtValue = fmt.Sprintf("%s: %s", key, fmtValue)
			} else {
				fmtValue = fmt.Sprintf("%s:\n  %s", key, indent(fmtValue))
			}
		case []interface{}:
			fmtValue = fmt.Sprintf("%s:\n  %s", key, indent(fmtValue))
		default:
			fmtValue = fmt.Sprintf("%s: %s", key, fmtValue)
		}

		parts[i] = fmtValue
	}

	joiner := "\n"

	if len(path) <= 1 {
		joiner = "\n\n"
	}

	return strings.Join(parts, joiner)
}

func formatYamlList(data []interface{}, path []string) string {
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
	if value, ok := data.(map[string]interface{}); ok {
		return formatYamlMap(value, path)
	}

	if value, ok := data.([]interface{}); ok {
		return formatYamlList(value, path)
	}

	if value, ok := data.(string); ok {
		return formatString(value)
	}

	return fmt.Sprint(data)
}

func Yaml(data interface{}) string {
	return yaml(data, make([]string, 0))
}
