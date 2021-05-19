#!/usr/bin/env bash

if [ $# -eq 0 ]; then
    echo "No arguments supplied"
    exit 2
fi

./extract-code-blocks.sh examples $1

for f in examples/$1/*/*.go ; do
    df=$(dirname $f)/$(basename $f .go)
    mkdir -p $df
    mv $f $df/main.go
done

ls -d ./examples/$1/*/* | xargs go build -o /dev/null -ldflags "-s -w"
if [ $? -ne 0 ]; then
    echo -e "Failed to build examples"
    exit 1
fi
