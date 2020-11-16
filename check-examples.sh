#!/usr/bin/env bash

./extract-code-blocks.sh

ls examples/v*/*/*.go | xargs -P $(($(nproc)*2)) -n 1 go build -o /dev/null
if [ $? -ne 0 ]; then
    echo -e "Failed to build examples"
    exit 1
fi
