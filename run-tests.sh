#!/usr/bin/env sh
set -e
golint *.go
go test -v
