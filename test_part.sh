#!/bin/bash

# Known paths that definitely won't validate correctly - mostly macros
declare -a EXCLUSIONS=(
    "./aws-cloudformation-templates/aws/services/CloudFormation/MacrosExamples/ShortHand/example.template"
    "./aws-cloudformation-templates/aws/solutions/StackSetsResource/TestResources/events-config.json"
    "./aws-cloudformation-templates/aws/services/CloudFormation/MacrosExamples/Count/test.yaml"
    "./aws-cloudformation-templates/aws/services/CloudFormation/MacrosExamples/Count/test_2.yaml"
    "./aws-cloudformation-templates/aws/services/CloudFormation/MacrosExamples/StringFunctions/string_example.yaml"
    "./aws-cloudformation-templates/aws/services/CloudFormation/MacrosExamples/PyPlate/python_example.yaml"
    "./aws-cloudformation-templates/aws/services/CloudFormation/MacrosExamples/Public-and-Private-Subnet-per-AZ/Create-Stack.yaml"
)

# Ignore exclusions
for path in "${EXCLUSIONS[@]}"; do
    if [ "$1" = "$path" ]; then
        exit
    fi
done

# Run the binary
errors=$(
    $CFN_FORMAT_TEST_BINARY --verify $1 2>&1 >/dev/null
    $CFN_FORMAT_TEST_BINARY --verify -j $1 2>&1 >/dev/null
)

# Print out any errors
if [ -n "$errors" ]; then
    echo ">>> $1"
    echo $errors
    echo
fi
