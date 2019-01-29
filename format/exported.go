/*
Package format provides functions for formatting CloudFormation templates
using an opinionated, idiomatic format as used in AWS documentation.

For each function, CloudFormation templates should be represented using
a map[string]interface{} as output by other libraries that parse JSON/YAML
such as github.com/awslabs/goformation and encoding/json.

Comments can be passed along with the template data in the following format:

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

Empty string keys are taken to represent a comment on the overall node
that the comment is attached to. Numeric keys can be used to reference
elements of arrays in the source data.
*/
package format

// Yaml formats the CloudFormation template as a Yaml string
func Yaml(data map[string]interface{}) string {
	return newFormatter("yaml", data).format()
}

// YamlWithComments formats the CloudFormation template
// as a Yaml string with comments as provided
func YamlWithComments(data map[string]interface{}, comments map[interface{}]interface{}) string {
	return newFormatterWithComments("yaml", data, comments).format()
}

// Json formats the CloudFormation template as a Json string
func Json(data map[string]interface{}) string {
	return newFormatter("json", data).format()
}

// JsonWithComments formats the CloudFormation template
// as a Json string with comments as provided
func JsonWithComments(data map[string]interface{}, comments map[interface{}]interface{}) string {
	return newFormatterWithComments("json", data, comments).format()
}
