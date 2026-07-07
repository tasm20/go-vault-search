#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")"

app_name="$(basename "$PWD")"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o "${app_name}_amd64_linux" .
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -trimpath -o "${app_name}_arm64_darwin" .
