#!/bin/bash

# Run the binary
errors=$(
    $CFN_FORMAT_TEST_BINARY $1 2>&1 >/dev/null
    $CFN_FORMAT_TEST_BINARY -j $1 2>&1 >/dev/null
)

# Print out any errors
if [ -n "$errors" ]; then
    echo ">>> $1"
    echo $errors
    echo
fi
