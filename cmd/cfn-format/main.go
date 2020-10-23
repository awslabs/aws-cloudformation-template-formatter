package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws-cloudformation/rain/cfn"
	"github.com/aws-cloudformation/rain/cfn/format"
	"github.com/aws-cloudformation/rain/cfn/parse"

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

var compactFlag bool
var jsonFlag bool
var verifyFlag bool
var writeFlag bool

func init() {
	pflag.BoolVarP(&compactFlag, "compact", "c", false, "Produce more compact output.")
	pflag.BoolVarP(&jsonFlag, "json", "j", false, "Output the template as JSON (default format: YAML).")
	pflag.BoolVarP(&verifyFlag, "verify", "v", false, "Check if the input is already correctly formatted and exit.\nThe exit status will be 0 if so and 1 if not.")
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
	var input []byte
	var source cfn.Template
	var err error

	pflag.Parse()
	args := pflag.Args()

	if len(args) == 1 {
		// Reading from a file
		fileName = args[0]
		input, err = ioutil.ReadFile(fileName)
		if err != nil {
			die(err.Error())
		}
	} else if !termutil.Isatty(os.Stdin.Fd()) {
		if writeFlag {
			// Can't use write without a filename!
			die("Can't write back to a file when reading from stdin")
		}

		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			die(err.Error())
		}
	} else {
		pflag.Usage()
	}

	source, err = parse.String(string(input))
	if err != nil {
		die(err.Error())
	}

	// Format the output
	options := format.Options{}

	if jsonFlag {
		options.Style = format.JSON
	}

	if compactFlag {
		options.Compact = true
	}

	output := format.Template(source, options)

	if verifyFlag {
		if len(fileName) > 0 {
			fmt.Fprint(os.Stderr, fileName+": ")
		}
		if strings.TrimSpace(string(input)) == output {
			fmt.Fprintln(os.Stderr, "formatted OK")
			os.Exit(0)
		} else {
			die("would reformat")
		}
	}

	// Verify the output is valid
	err = parse.Verify(source, output)
	if err != nil {
		die(err.Error())
	}

	if writeFlag {
		ioutil.WriteFile(fileName, []byte(output), 0644)
	} else {
		fmt.Println(output)
	}
}
