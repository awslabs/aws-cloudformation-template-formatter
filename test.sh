#!/bin/bash

function verify {
    go run main.go --verify $1 >/dev/null
    go run main.go --verify -j $1 >/dev/null
}

export -f verify

git clone --depth 1 https://github.com/awslabs/aws-cloudformation-templates
find ./aws-cloudformation-templates -iname *.template -exec bash -c 'verify "$0"' {} \;
rm -r aws-cloudformation-templates
