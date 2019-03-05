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
  -j      Output the template as JSON (default format: YAML).
  --help  Show this message and exit.
```

### Go package documentation

```go
import "github.com/awslabs/aws-cloudformation-template-formatter/format"
```

Package format provides functions for formatting CloudFormation
templates using an opinionated, idiomatic format as used in AWS
documentation.

For each function, CloudFormation templates should be represented using
a `map[string]interface{}` as output by other libraries that parse
JSON/YAML such as `github.com/awslabs/goformation` and `encoding/json`.

Comments can be passed along with the template data in the following
format:

```go
map[interface{}]interface{}{
    "": "This is a top-level comment",
    "Resources": map[interface{}]interface{}{
        "": "This is a comment on the whole `Resources` property",
        "MyBucket": map[interface{}]interface{}{
            "Properties": map[interface{}]interface{}{
                "BucketName": "This is a comment on BucketName",
            },
        },
    },
}
```

Empty string keys are taken to represent a comment on the overall node
that the comment is attached to. Numeric keys can be used to reference
elements of arrays in the source data.

#### Functions

```go
func Json(data map[string]interface{}) string
    Json formats the CloudFormation template as a Json string
```

```go
func JsonWithComments(data map[string]interface{}, comments map[interface{}]interface{}) string
    JsonWithComments formats the CloudFormation template as a Json string
    with comments as provided
```

```go
func Yaml(data map[string]interface{}) string
    Yaml formats the CloudFormation template as a Yaml string
```

```go
func YamlWithComments(data map[string]interface{}, comments map[interface{}]interface{}) string
    YamlWithComments formats the CloudFormation template as a Yaml string
    with comments as provided
```
