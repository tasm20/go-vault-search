#!/bin/bash
go install ./
GOOS=linux GOARCH=amd64 go build -o "$(basename "$(pwd)")"_amd64_linux
GOOS=darwin GOARCH=arm64 go build -o "$(basename "$(pwd)")"_arm64_darwin
