#!/usr/bin/env sh
set -e
golint ./...
go test -coverpkg=$(go list ./... | grep -v /test | paste -sd "," -) -coverprofile=coverage.txt -covermode=atomic -v $(go list ./...)
