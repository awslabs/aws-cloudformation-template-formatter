# To do

## Refactoring

* Drop the dependency on goformation and use sanathkr/yaml directly

    This is because we don't want to necessarily validate the template.
    Some edge cases using intrinsics and not processing them cause failures.

* Create a Template model with functions for ToJson and ToYaml

* Change the output from string to []byte

* Make `common.go` a bit more intuitive to maintain

## Improvements

* Order resources based on dependencies between them

* Add long/multi-line string handling

* deal with ordering in more resource types

    * e.g. IAM policies

## Features

* vim plugin

* VSCode plugin

* Use a comment-preserving JSON/YAML parser
