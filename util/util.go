package util

import (
	"fmt"

	"github.com/awslabs/goformation/cloudformation"
	"github.com/awslabs/goformation/intrinsics"
	"github.com/google/go-cmp/cmp"

	"encoding/json"

	"github.com/awslabs/goformation"
)

func read(source *cloudformation.Template) (map[string]interface{}, error) {
	// Convert to JSON and back just to get rid of the goformation types
	sourceJson, err := json.Marshal(source)
	if err != nil {
		return nil, fmt.Errorf("Internal error: %s", err.Error())
	}

	var output map[string]interface{}
	err = json.Unmarshal(sourceJson, &output)
	if err != nil {
		return nil, fmt.Errorf("Internal error: %s", err.Error())
	}

	return output, nil
}

func ReadFile(fileName string) (map[string]interface{}, error) {
	source, err := goformation.OpenWithOptions(fileName, &intrinsics.ProcessorOptions{
		NoProcess: true,
	})
	if err != nil {
		return nil, fmt.Errorf("Unable to process input template (%s): %s", fileName, err.Error())
	}

	return read(source)
}

func ReadString(input string) (map[string]interface{}, error) {
	source, err := goformation.ParseYAMLWithOptions([]byte(input), &intrinsics.ProcessorOptions{
		NoProcess: true,
	})
	if err != nil {
		return nil, fmt.Errorf("Unable to process input template: %s", err.Error())
	}

	return read(source)
}

func VerifyOutput(source map[string]interface{}, output string) error {
	// Check it matches the original
	validate, err := ReadString(output)
	if err != nil {
		return err
	}

	if diff := cmp.Diff(source, validate); diff != "" {
		return fmt.Errorf("Semantic difference after formatting:\n%s", diff)
	}

	return nil
}
