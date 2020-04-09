#!/usr/bin/env sh
set -e
echo "Ensuring all dependencies are present in LICENSE-3rdparty.csv ..."
go mod tidy
ALL_DEPS=`cat go.sum | awk '{print $1}' | uniq | sort | sed "s|^\(.*\)|go.sum,\1,|"`
DEPS_NOT_FOUND=""
for one_dep in `echo $ALL_DEPS`; do
    set +e
    cat LICENSE-3rdparty.csv | grep "$one_dep" > /dev/null 2>&1
    if [ $? -ne 0 ]; then
        DEPS_NOT_FOUND="${DEPS_NOT_FOUND}\n${one_dep}<LICENSE>,<COPYRIGHT>"
    fi
    set -e
done
if [ -n "$DEPS_NOT_FOUND" ]; then
    echo "Some dependencies were not found in LICENSE-3rdparty.csv, please add: $DEPS_NOT_FOUND"
    exit 1
else
    echo "LICENSE-3rdparty.csv is up to date"
fi

# make sure the below installed dependencies don't get added to go.mod/go.sum
# unfortunately there's no better way to fix this than change directory
# this might get solved in Go 1.14: https://github.com/golang/go/issues/30515
pushd /tmp
go get -u golang.org/x/lint/golint
go get -u gotest.tools/gotestsum@v0.4.1
popd

golint ./...
gotestsum --format testname -- -coverpkg=$(go list ./... | grep -v /test | paste -sd "," -) -coverprofile=coverage.txt -covermode=atomic -v $(go list ./...)
go mod tidy
