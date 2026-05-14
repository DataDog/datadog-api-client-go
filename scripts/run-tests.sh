#!/usr/bin/env bash
set -e

# make sure the below installed dependencies don't get added to go.mod/go.sum
# unfortunately there's no better way to fix this than change directory
# this might get solved in Go 1.14: https://github.com/golang/go/issues/30515
echo "Installing gotestsum"
cd "$(mktemp -d)"
GO111MODULE=on go install gotest.tools/gotestsum@latest
cd -

echo "Running mod tidy and cleanup test cache"
go mod tidy
go clean -testcache

# Run the same in tests submodule
cd tests
echo "Running mod tidy and cleanup test cache in test module"
go mod tidy
go clean -testcache

if [ "$RECORD" == "none" ]; then
    echo "Running with no RECORD"
    gotestsum --rerun-fails=1 --format short-verbose --packages ./... -- -timeout=20m $TESTARGS
else
    echo "Running with RECORD"
    gotestsum --format short-verbose --packages ./... -- $TESTARGS
fi

cd ..
