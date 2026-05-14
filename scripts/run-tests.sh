#!/usr/bin/env bash
set -e

GOTESTSUM_VERSION="1.13.0"

if ! command -v gotestsum &> /dev/null; then
    echo "Installing gotestsum v${GOTESTSUM_VERSION} (pre-built binary)"
    curl -sSfL "https://github.com/gotestyourself/gotestsum/releases/download/v${GOTESTSUM_VERSION}/gotestsum_${GOTESTSUM_VERSION}_linux_amd64.tar.gz" \
        | tar -xz -C "$(go env GOPATH)/bin" gotestsum
else
    echo "gotestsum already installed"
fi

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
