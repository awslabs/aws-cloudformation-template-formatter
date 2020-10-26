#!/bin/bash

# This script will build rain for all platforms

set -e

NAME=cfn-format
OUTPUT_DIR=dist

VERSION=$(git describe --abbrev=0)

declare -A PLATFORMS=([linux]=linux [darwin]=macos [windows]=windows)
declare -A ARCHITECTURES=([386]=i386 [amd64]=amd64 [arm]=arm [arm64]=arm64)

# Run tests first
golint -set_exit_status ./... || exit 1
go vet ./... || exit 1
go test ./... || exit 1

echo "Building $NAME $VERSION..."

for platform in ${!PLATFORMS[@]}; do
    for architecture in ${!ARCHITECTURES[@]}; do
        if [[ "$architecture" == arm* && "$platform" != "linux" ]]; then
            continue
        fi

        if [[ "$architecture" == "386" && "$platform" == "darwin" ]]; then
            continue
        fi

        echo "$platform/$architecture..."

        full_name="${NAME}-${VERSION}_${PLATFORMS[$platform]}-${ARCHITECTURES[$architecture]}"
        bin_name="$NAME"

        if [ "$platform" == "windows" ]; then
            bin_name="${NAME}.exe"
        fi

        mkdir -p "$OUTPUT_DIR/$full_name"

        eval GOOS=$platform GOARCH=$architecture go build -o "$OUTPUT_DIR/${full_name}/${bin_name}" "./cmd/cfn-format/*"
        cp LICENSE "$OUTPUT_DIR/$full_name"
        cp README.md "$OUTPUT_DIR/$full_name"

        cd "$OUTPUT_DIR"
        zip -9 -r "${full_name}.zip" "$full_name"
        rm -r "$full_name"
        cd -
    done
done

echo "All done."
