[![GitHub version](https://badge.fury.io/gh/awslabs%2Faws-cloudformation-template-formatter.svg)](https://badge.fury.io/gh/awslabs%2Faws-cloudformation-template-formatter)

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/cfn-format)

# AWS CloudFormation Template Formatter

This repository contains `cfn-format`, a command line tool and Go library that reads in an existing [AWS CloudFormation](https://aws.amazon.com/cloudformation/) template and outputs a cleanly-formatted, easy-to-read copy of the same template adhering to standards as used in AWS documentation. `cfn-format` can output either YAML or JSON as desired.

## License

This project is licensed under the Apache 2.0 License. 

## Usage

To use `cfn-format`, you can clone this repository and run `go build` or download the [latest release](https://github.com/awslabs/aws-cloudformation-template-formatter/releases/latest) for your operating system.

If you're using [vim](https://www.vim.org/), you can add the following to your `.vimrc` to automate running `cfn-format` when you save a `.template` file:

```vim
autocmd BufWritePost *.template silent !cfn-format -w % 2>/dev/null
```

### Command-line tool

```console
Usage: cfn-format [OPTION...] [FILENAME]

  AWS CloudFormation Format is a tool that reads a CloudFormation template
  and outputs the same template, formatted according to the same standards
  used in AWS documentation.

  If FILENAME is not supplied, cfn-format will read from STDIN.

Options:
  --help    Show this message and exit.
  -j, --json    Output the template as JSON (default format: YAML).
  -w, --write   Write the output back to the file rather than to stdout.
```

### Go package documentation

**The API for cfn-format will be changing in 1.0.0**

To see the current Go documentation for cfn-format, please check [format/exported.go](https://github.com/awslabs/aws-cloudformation-template-formatter/blob/0.3.0/format/exported.go).
