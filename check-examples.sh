#!/usr/bin/env bash

./extract-code-blocks.sh examples

for f in examples/*/*/*.go ; do
    df=$(dirname $f)/$(basename $f .go)
    mkdir -p $df
    mv $f $df/main.go
done

ls -d ./examples/*/*/* | xargs go build -o /dev/null -ldflags "-s -w"
if [ $? -ne 0 ]; then
    echo -e "Failed to build examples"
    exit 1
fi
