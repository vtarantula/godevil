#!/usr/bin/env bash
# Remove the CGO comments when the C library is built
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/godevil_linux_amd64 cmd/main.go
env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o build/godevil_windows_amd64 cmd/main.go
