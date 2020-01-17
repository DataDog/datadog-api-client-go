#!/usr/bin/env sh
set -e
golint *.go
go test -coverprofile=coverage.txt -covermode=atomic -v
