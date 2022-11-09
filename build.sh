#!/usr/bin/env bash
env GOOS=linux GOARCH=amd64 go build -o build/godevil_linux_amd64 cmd/main.go
env GOOS=windows GOARCH=amd64 go build -o build/godevil_windows_amd64 cmd/main.go
