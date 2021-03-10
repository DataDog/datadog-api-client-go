#!/usr/bin/env bash
# make sure the below installed dependencies don't get added to go.mod/go.sum
# unfortunately there's no better way to fix this than change directory
# this might get solved in Go 1.14: https://github.com/golang/go/issues/30515
cd `mktemp -d`
GO111MODULE=on go get -u golang.org/x/lint/golint
GO111MODULE=on go get -u gotest.tools/gotestsum
cd -

golint ./...
go mod tidy
go clean -testcache
gotestsum --format short-verbose --rerun-fails --rerun-fails-max-failures=20000 --raw-command -- ./run-go-tests.sh
RESULT+=$?

# Always run integration-only scenarios
set -e
if [ "$RECORD" != "none" -a -n "$DD_TEST_CLIENT_API_KEY" -a -n "$DD_TEST_CLIENT_APP_KEY" ]; then
  BDD_TAGS="@integration-only" RECORD=none gotestsum --format short-verbose --rerun-fails --rerun-fails-max-failures=20000 --raw-command -- ./run-go-tests.sh
fi

exit $RESULT
