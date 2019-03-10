package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/awslabs/aws-cloudformation-template-formatter/format"
	"github.com/awslabs/aws-cloudformation-template-formatter/util"
)

var usage = `Usage: cfn-format [-j] <filename>

  AWS CloudFormation Format is a tool that reads a CloudFormation template
  and outputs the same template, formatted according to the same standards
  used in AWS documentation.

Options:
  -j        Output the template as JSON (default format: YAML).
  -w        Write the output back to the file rather than to stdout.
  --verify  Verify that the formatted output is semantically identical to the input.
            Use this option to guarantee your template has not changed.
  --help    Show this message and exit.
`

func die(message string) {
	fmt.Fprintf(os.Stderr, message)
	os.Exit(1)
}

func help() {
	die(usage)
}

func main() {
	style := "yaml"
	verify := false
	write := false

	// Parse options
	if len(os.Args) < 2 {
		help()
	}

	for _, arg := range os.Args[1 : len(os.Args)-1] {
		switch arg {
		case "-j":
			style = "json"
		case "--verify":
			verify = true
		case "-w":
			write = true
		case "-h", "--help":
			help()
		}
	}

	// Get the filename
	fileName := os.Args[len(os.Args)-1]
	if fileName == "--help" {
		help()
	}

	// Read the source template
	source, err := util.ReadFile(fileName)
	if err != nil {
		die(err.Error())
	}

	// Format the output
	var output string
	if style == "json" {
		output = format.Json(source)
	} else {
		output = format.Yaml(source)
	}

	if verify {
		err := util.VerifyOutput(source, output)
		if err != nil {
			die(err.Error())
		}
	}

	if write {
		ioutil.WriteFile(fileName, []byte(output), 0644)
	} else {
		fmt.Println(output)
	}
}
