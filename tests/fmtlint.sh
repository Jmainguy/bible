#!/bin/bash
FormatCheck=$(gofmt -l *.go | wc -l)
if [ $FormatCheck -gt 0 ]; then
    gofmt -l *.go
    echo "gofmt -w *.go your code please."
    exit 1
fi

## Run staticcheck
staticcheck ./...
if [ $? -gt 0 ]; then
    exit 1
fi

## Run go vet
go vet ./...
if [ $? -gt 0 ]; then
    exit 1
fi
