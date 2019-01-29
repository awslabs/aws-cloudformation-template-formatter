# AWS CloudFormation Template Formatter

This repository contains `cfn-format`, a command line tool and Go library that reads in an existing [AWS CloudFormation](https://aws.amazon.com/cloudformation/) template and outputs a cleanly-formatted, easy-to-read copy of the same template adhering to standards as used in AWS documentation. `cfn-format` can output either YAML or JSON as desired.

## License

This project is licensed under the Apache 2.0 License. 

## Usage

### Command-line tool

```console
Usage: cfn-format [-j] <filename>

  AWS CloudFormation Format is a tool that reads a CloudFormation template
  and outputs the same template, formatted according to the same standards
  used in AWS documentation.

Options:
  -j      Output the template as JSON (default format: YAML).
  --help  Show this message and exit.
```

### Go library

To use the Go library, import the `format` package and then use one of the following exported functions:
package format // import "codecommit/builders/cfn-format/format"

func Json(data map[string]interface{}) string
func JsonWithComments(data interface{}, comments map[interface{}]interface{}) string
func Yaml(data map[string]interface{}) string
func YamlWithComments(data interface{}, comments map[interface{}]interface{}) string
type Struct struct{ ... }
