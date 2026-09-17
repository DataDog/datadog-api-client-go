#!/usr/bin/env bash
set -e

echo "Ensuring all dependencies are present in LICENSE-3rdparty.csv ..."
go mod tidy
ALL_DEPS=$(cat go.sum tests/go.sum | awk '{print $1}' | uniq | sort | sed "s|^\(.*\)|go.sum,\1,|")
DEPS_NOT_FOUND=""
set +e
for one_dep in $ALL_DEPS; do
    cat LICENSE-3rdparty.csv | grep "$one_dep" > /dev/null 2>&1
    if [ $? -ne 0 ]; then
        DEPS_NOT_FOUND="${DEPS_NOT_FOUND}\n${one_dep}<LICENSE>,<COPYRIGHT>"
    fi
done
set -e
if [ -n "$DEPS_NOT_FOUND" ]; then
    echo "Some dependencies were not found in LICENSE-3rdparty.csv, please add: $DEPS_NOT_FOUND"
    exit 1
else
    echo "LICENSE-3rdparty.csv is up to date"
fi

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
