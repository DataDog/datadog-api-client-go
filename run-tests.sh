#!/usr/bin/env sh
set -e
golint ./...
gotestsum --format testname -- -coverpkg=$(go list ./... | grep -v /test | paste -sd "," -) -coverprofile=coverage.txt -covermode=atomic -v $(go list ./...)
