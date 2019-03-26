package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/awslabs/aws-cloudformation-template-formatter/format"
	"github.com/awslabs/aws-cloudformation-template-formatter/parse"

	"github.com/andrew-d/go-termutil"
	"github.com/spf13/pflag"
)

var usage = `Usage: cfn-format [OPTION...] [FILENAME]

  AWS CloudFormation Format is a tool that reads a CloudFormation template
  and outputs the same template, formatted according to the same standards
  used in AWS documentation.

  If FILENAME is not supplied, cfn-format will read from STDIN.

Options:
  --help    Show this message and exit.`

var jsonFlag bool
var writeFlag bool

func init() {
	pflag.BoolVarP(&jsonFlag, "json", "j", false, "Output the template as JSON (default format: YAML).")
	pflag.BoolVarP(&writeFlag, "write", "w", false, "Write the output back to the file rather than to stdout.")

	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		pflag.PrintDefaults()
		os.Exit(1)
	}
}

func die(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func main() {
	var fileName string
	var source map[string]interface{}
	var err error

	pflag.Parse()
	args := pflag.Args()

	if len(args) == 1 {
		// Reading from a file
		fileName = args[0]
		source, err = parse.ReadFile(fileName)
		if err != nil {
			die(err.Error())
		}
	} else if !termutil.Isatty(os.Stdin.Fd()) {
		if writeFlag {
			// Can't use write without a filename!
			die("Can't write back to a file when reading from stdin")
		}

		source, err = parse.Read(os.Stdin)
		if err != nil {
			die(err.Error())
		}
	} else {
		pflag.Usage()
	}

	// Format the output
	var output string
	if jsonFlag {
		output = format.Json(source)
	} else {
		output = format.Yaml(source)
	}

	// Verify the output is valid
	err = parse.VerifyOutput(source, output)
	if err != nil {
		die(err.Error())
	}

	if writeFlag {
		ioutil.WriteFile(fileName, []byte(output), 0644)
	} else {
		fmt.Println(output)
	}
}
