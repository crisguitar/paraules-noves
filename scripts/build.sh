#!/usr/bin/env bash

rm -rf build

GOOS=linux GOARCH=amd64 go build -o build/app cmd/main.go

echo "Done..."
