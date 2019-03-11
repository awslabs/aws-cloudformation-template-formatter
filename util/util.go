package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/google/go-cmp/cmp"
	yaml "github.com/sanathkr/go-yaml"
	yamlwrapper "github.com/sanathkr/yaml"
)

var allTags = []string{
	"Ref", "GetAtt", "Base64", "FindInMap", "GetAZs",
	"ImportValue", "Join", "Select", "Split", "Sub",
	"Equals", "Cidr", "And", "If", "Not", "Or",
}

type tagUnmarshalerType struct {
}

var tagUnmarshaler = &tagUnmarshalerType{}

func init() {
	for _, tag := range allTags {
		yaml.RegisterTagUnmarshaler("!"+tag, tagUnmarshaler)
	}
}

func (t *tagUnmarshalerType) UnmarshalYAMLTag(tag string, value reflect.Value) reflect.Value {
	prefix := "Fn::"
	if tag == "Ref" || tag == "Condition" {
		prefix = ""
	}
	tag = prefix + tag

	output := reflect.ValueOf(make(map[interface{}]interface{}))
	key := reflect.ValueOf(tag)
	output.SetMapIndex(key, value)

	return output
}

func ReadFile(fileName string) (map[string]interface{}, error) {
	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to read file: %s", err)
	}

	return ReadBytes(source)
}

func ReadBytes(input []byte) (map[string]interface{}, error) {
	parsed, err := yamlwrapper.YAMLToJSON(input)
	if err != nil {
		return nil, fmt.Errorf("Invalid YAML: %s", err)
	}

	var output map[string]interface{}
	err = json.Unmarshal(parsed, &output)
	if err != nil {
		return nil, fmt.Errorf("Invalid YAML: %s", err)
	}

	return output, nil
}

func VerifyOutput(source map[string]interface{}, output []byte) error {
	// Check it matches the original
	validate, err := ReadBytes(output)
	if err != nil {
		return err
	}

	if diff := cmp.Diff(source, validate); diff != "" {
		return fmt.Errorf("Semantic difference after formatting:\n%s", diff)
	}

	return nil
}
