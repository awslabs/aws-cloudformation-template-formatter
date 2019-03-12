[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/cfn-format)

# AWS CloudFormation Template Formatter

This repository contains `cfn-format`, a command line tool and Go library that reads in an existing [AWS CloudFormation](https://aws.amazon.com/cloudformation/) template and outputs a cleanly-formatted, easy-to-read copy of the same template adhering to standards as used in AWS documentation. `cfn-format` can output either YAML or JSON as desired.

## License

This project is licensed under the Apache 2.0 License. 

## Usage

To use `cfn-format`, you can clone this repository and run `go build` or download the [latest release](https://github.com/awslabs/aws-cloudformation-template-formatter/releases/latest) for your operating system.

### Command-line tool

```console
Usage: cfn-format [-j] <filename>

  AWS CloudFormation Format is a tool that reads a CloudFormation template
  and outputs the same template, formatted according to the same standards
  used in AWS documentation.

Options:
  -j        Output the template as JSON (default format: YAML).
  -w        Write the output back to the file rather than to stdout.
  --help    Show this message and exit.
```

### Go package documentation

**The API for cfn-format will be changing in 1.0.0**

To see the current Go documentation for cfn-format, please check <format/exported.go>.
