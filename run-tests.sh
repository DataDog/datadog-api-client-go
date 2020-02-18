#!/usr/bin/env sh
set -e
golint ./...
go test -coverprofile=coverage.txt -covermode=atomic -v $(go list ./...)
