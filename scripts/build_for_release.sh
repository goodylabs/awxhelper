#!/bin/bash

set -e

VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
echo "Building version: $VERSION"

mkdir -p bin

ldflags="-X 'github.com/goodylabs/awxhelper/cmd.version=$VERSION'"

GOOS=linux GOARCH=amd64 go build -ldflags "$ldflags" -o bin/linux-amd64
GOOS=linux GOARCH=arm64 go build -ldflags "$ldflags" -o bin/linux-arm64
GOOS=darwin GOARCH=amd64 go build -ldflags "$ldflags" -o bin/darwin-amd64
GOOS=darwin GOARCH=arm64 go build -ldflags "$ldflags" -o bin/darwin-arm64
