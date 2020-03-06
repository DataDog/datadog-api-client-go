#!/usr/bin/env sh
set -e
golint ./...
go test -coverpkg=github.com/DataDog/datadog-api-client-go/api/v1/datadog -coverprofile=coverage.txt -covermode=atomic -v $(go list ./...)
