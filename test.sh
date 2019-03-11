#!/bin/bash

# Build a test binary
export CFN_FORMAT_TEST_BINARY=$(mktemp)
go build -o $CFN_FORMAT_TEST_BINARY

# Get the example templates
git clone --depth 1 https://github.com/awslabs/aws-cloudformation-templates &>/dev/null

# Run the tests
find ./aws-cloudformation-templates -iname "*.template" -exec ./test_part.sh {} \;
find ./aws-cloudformation-templates -iname "*.json" -exec ./test_part.sh {} \;
find ./aws-cloudformation-templates -iname "*.yaml" -exec ./test_part.sh {} \;

# CLean up
rm -rf aws-cloudformation-templates
rm $CFN_FORMAT_TEST_BINARY
