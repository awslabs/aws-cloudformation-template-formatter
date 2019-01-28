package format

import (
	"fmt"
	"reflect"
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

func sortKeys(data map[string]interface{}, path []interface{}) []string {
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

func intrinsicKey(data map[string]interface{}) (string, bool) {
	if len(data) != 1 {
		return "", false
	}

	// We know there's one key
	key := reflect.ValueOf(data).MapKeys()[0].String()
	if key == "Ref" || strings.HasPrefix(key, "Fn::") {
		return key, true
	}

	return "", false
}

func formatString(data string) string {
	quote := false

	switch {
	case strings.ContainsAny(data, "\n"),
		data == "",
		data == "Yes" || data == "No",
		strings.ContainsAny(string(data[0]), "0123456789!&*?,#|>@`\"'[{: -\\"),
		strings.ContainsAny(data, "`\"'"):
		quote = true
	}

	if quote {
		return fmt.Sprintf("%q", data)
	}

	return data
}

func getFromPath(data map[string]interface{}, path []interface{}) interface{} {
	value := interface{}(data)

	for _, part := range path {
		switch v := value.(type) {
		case map[string]interface{}:
			value = v[part.(string)]
		case []interface{}:
			value = v[part.(int)]
		default:
			panic(fmt.Sprintf("Something has gone wrong with the path: %v\n", path))
		}
	}

	return value
}
