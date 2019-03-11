package util

import (
	"fmt"
	"io/ioutil"

	"github.com/google/go-cmp/cmp"
	"github.com/sanathkr/yaml"
)

func ReadFile(fileName string) (map[string]interface{}, error) {
	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return ReadBytes(source)
}

func ReadBytes(input []byte) (map[string]interface{}, error) {
	var output map[string]interface{}
	err := yaml.Unmarshal(input, &output)
	if err != nil {
		return nil, err
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
