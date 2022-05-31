#!/usr/bin/env bash

# make sure the below installed dependencies don't get added to go.mod/go.sum
# unfortunately there's no better way to fix this than change directory
# this might get solved in Go 1.14: https://github.com/golang/go/issues/30515
cd "$(mktemp -d)"
GO111MODULE=on go install golang.org/x/lint/golint@latest
GO111MODULE=on go install gotest.tools/gotestsum@latest
cd -

golint ./...
go mod tidy
go clean -testcache

# Run the same in tests submodule
cd tests
golint ./...
go mod tidy
go clean -testcache
gotestsum --format short-verbose ./...
cd ..
