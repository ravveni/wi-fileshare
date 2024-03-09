#!/bin/sh

BINARY_NAME="wi-fileshare"

# linux
GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64-${BINARY_NAME}
GOOS=linux GOARCH=arm64 go build -o bin/linux-arm64-${BINARY_NAME}

# mac
GOOS=darwin GOARCH=amd64 go build -o bin/darwin-amd64-${BINARY_NAME}
GOOS=darwin GOARCH=arm64 go build -o bin/darwin-arm64-${BINARY_NAME}

# windows
GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64-${BINARY_NAME}.exe
