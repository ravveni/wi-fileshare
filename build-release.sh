#!/bin/sh

BINARY_NAME="wi-fileshare"

# linux
GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-amd64-linux
GOOS=linux GOARCH=arm64 go build -o bin/${BINARY_NAME}-arm64-linux

# mac
GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-amd64-darwin
GOOS=darwin GOARCH=arm64 go build -o bin/${BINARY_NAME}-arm64-darwin

# windows
GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-amd64.exe
