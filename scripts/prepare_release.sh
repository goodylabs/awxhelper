#!/bin/bash

mkdir -p bin

GOOS=linux GOARCH=amd64 go build -ldflags "$ldflags" -o bin/awxhelper-linux-amd64
GOOS=linux GOARCH=arm64 go build -ldflags "$ldflags" -o bin/awxhelper-linux-arm64
GOOS=darwin GOARCH=amd64 go build -ldflags "$ldflags" -o bin/awxhelper-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -ldflags "$ldflags" -o bin/awxhelper-darwin-arm64
