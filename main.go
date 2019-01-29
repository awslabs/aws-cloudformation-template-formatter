package main

import (
	"fmt"
	"os"

	"./format"

	"encoding/json"

	"github.com/awslabs/goformation"
	"github.com/awslabs/goformation/intrinsics"
)

var usage = `Usage: cfn-format [-j] <filename>

  AWS CloudFormation Format is a tool that reads a CloudFormation template
  and outputs the same template, formatted according to the same standards
  used in AWS documentation.

Options:
  -j      Output the template as JSON (default format: YAML).
  --help  Show this message and exit.
`

func die() {
	fmt.Fprintf(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	style := "yaml"
	if len(os.Args) == 3 {
		if os.Args[1] == "-j" {
			style = "json"
		} else {
			die()
		}
	} else if len(os.Args) != 2 {
		die()
	}

	fileName := os.Args[len(os.Args)-1]

	if fileName == "--help" {
		die()
	}

	// We're literally just using this to parse the JSON/YAML
	source, err := goformation.OpenWithOptions(fileName, &intrinsics.ProcessorOptions{
		NoProcess: true,
	})
	if err != nil {
		panic("Unable to process input template: " + err.Error())
	}

	// Convert to JSON and back just to get rid of the goformation types
	sourceJson, err := json.Marshal(source)
	sourceValue := make(map[string]interface{})
	err = json.Unmarshal(sourceJson, &sourceValue)

	// YAMLise!
	if style == "json" {
		fmt.Println(format.Json(sourceValue))
	} else {
		fmt.Println(format.Yaml(sourceValue))
	}
}
