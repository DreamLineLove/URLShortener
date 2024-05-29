#!/bin/bash

GOOS=windows GOARCH=amd64 go build -o bin/windows.exe .
GOOS=linux GOARCH=amd64 go build -o bin/linux .
GOOS=darwin GOARCH=amd64 go build -o bin/intel-mac .
GOOS=darwin GOARCH=arm64 go build -o bin/apple-silicon-mac .
