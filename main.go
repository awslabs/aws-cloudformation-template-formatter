package main

//go:generate ./generate/generate.sh

import (
	"codecommit/builders/cfn-format/format"
	"fmt"
	"os"

	"encoding/json"

	"github.com/awslabs/goformation"
	"github.com/awslabs/goformation/intrinsics"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}

	fileName := os.Args[1]

	// We're literally just using this to parse the JSON/YAML
	source, err := goformation.OpenWithOptions(fileName, &intrinsics.ProcessorOptions{
		NoProcess: true,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Convert to JSON and back just to get rid of the goformation types
	sourceJson, err := json.Marshal(source)
	sourceValue := make(map[string]interface{})
	err = json.Unmarshal(sourceJson, &sourceValue)

	// YAMLise!
	fmt.Println(format.Yaml(sourceValue))
}
