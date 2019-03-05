#!/bin/bash

# This script will build cfn-format for all platforms

# Run tests first

go vet ./... || exit 1

go test ./... || exit 1

declare -A platforms=([linux]=linux [darwin]=osx [windows]=windows)
declare -A architectures=([386]=i386 [amd64]=amd64)

DESTDIR=dist

mkdir -p $DESTDIR

echo "Building cfn-format"

for platform in ${!platforms[@]}; do
    for architecture in ${!architectures[@]}; do
        echo "... $platform $architecture..."

        name=cfn-format-${platforms[$platform]}-${architectures[$architecture]}

        if [ "$platform" == "windows" ]; then
            name=${name}.exe
        fi

        GOOS=$platform GOARCH=$architecture go build -o $DESTDIR/$name
    done
done

echo "All done."
