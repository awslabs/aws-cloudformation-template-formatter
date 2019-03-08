package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/awslabs/aws-cloudformation-template-formatter/format"
	"github.com/awslabs/goformation/intrinsics"
	"github.com/google/go-cmp/cmp"

	"encoding/json"

	"github.com/awslabs/goformation"
)

var usage = `Usage: cfn-format [-j] <filename>

  AWS CloudFormation Format is a tool that reads a CloudFormation template
  and outputs the same template, formatted according to the same standards
  used in AWS documentation.

Options:
  -j        Output the template as JSON (default format: YAML).
  --verify  Verify that the formatted output is semantically identical to the input.
  --help    Show this message and exit.
`

func die() {
	fmt.Fprintf(os.Stderr, usage)
	os.Exit(1)
}

func read(fileName string) (output map[string]interface{}) {
	source, err := goformation.OpenWithOptions(fileName, &intrinsics.ProcessorOptions{
		NoProcess: true,
	})
	if err != nil {
		panic("Unable to process input template (" + fileName + "): " + err.Error())
	}

	// Convert to JSON and back just to get rid of the goformation types
	sourceJson, err := json.Marshal(source)
	if err != nil {
		panic("Internal error: " + err.Error())
	}

	err = json.Unmarshal(sourceJson, &output)
	if err != nil {
		panic("Internal error: " + err.Error())
	}

	return
}

func main() {
	style := "yaml"
	verify := false

	if len(os.Args) < 2 {
		die()
	}

	for _, arg := range os.Args[1 : len(os.Args)-1] {
		switch arg {
		case "-j":
			style = "json"
		case "--verify":
			verify = true
		case "-h", "--help":
			die()
		}
	}

	fileName := os.Args[len(os.Args)-1]

	if fileName == "--help" {
		die()
	}

	var output string

	source := read(fileName)

	if style == "json" {
		output = format.Json(source)
	} else {
		output = format.Yaml(source)
	}

	fmt.Println(output)

	if verify {
		// Write the output to a temporary file
		f, _ := ioutil.TempFile("", "*."+style)
		f.Write([]byte(output))
		f.Close()
		defer os.Remove(f.Name())

		// Check it matches the original
		if diff := cmp.Diff(source, read(f.Name())); diff != "" {
			fmt.Fprintln(os.Stderr, "Semantic difference after formatting "+fileName+" as "+style+":")
			fmt.Fprintln(os.Stderr, diff)
		}
	}
}
