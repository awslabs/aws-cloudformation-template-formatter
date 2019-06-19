[![GitHub version](https://badge.fury.io/gh/awslabs%2Faws-cloudformation-template-formatter.svg)](https://badge.fury.io/gh/awslabs%2Faws-cloudformation-template-formatter)
[![Snap Status](https://build.snapcraft.io/badge/awslabs/aws-cloudformation-template-formatter.svg)](https://build.snapcraft.io/user/awslabs/aws-cloudformation-template-formatter)

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/cfn-format)

# AWS CloudFormation Template Formatter

This repository contains `cfn-format`, a command line tool that reads in an existing [AWS CloudFormation](https://aws.amazon.com/cloudformation/) template and outputs a cleanly-formatted, easy-to-read copy of the same template adhering to standards as used in AWS documentation. `cfn-format` can output either YAML or JSON as desired.

## License

This project is licensed under the Apache 2.0 License. 

## Installation

You can install `cfn-format` in one of the following three ways:

* Use the [snap package](https://snapcraft.io/cfn-format)

* Download the [latest release](https://github.com/awslabs/aws-cloudformation-template-formatter/releases/latest) for your operating system.

* If you have [Go](https://golang.org/) (v1.12 or higher) installed, run the following:

    `GO111MODULE=on go get github.com/awslabs/aws-cloudformation-template-formatter/cmd/cfn-format`

## Usage

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
  -c, --compact   Produce more compact output.
  -j, --json      Output the template as JSON (default format: YAML).
  -v, --verify    Check if the input is already correctly formatted and exit.
                  The exit status will be 0 if so and 1 if not.
  -w, --write     Write the output back to the file rather than to stdout.
```

### Go package documentation

The `parse` and `format` packages have moved to become part of [rain](https://github.com/aws-cloudformation/rain).

To see the current Go documentation, please check [the Rain source code](https://github.com/aws-cloudformation/rain/tree/master/format/format.go).
