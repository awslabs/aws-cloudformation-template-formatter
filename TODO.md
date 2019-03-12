# To do

## Refactoring

* Pull out the yaml parsing into cfn-yaml

    Speak to Paul about this idea
    Submit a PR to goformation to use cfn-yaml

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
