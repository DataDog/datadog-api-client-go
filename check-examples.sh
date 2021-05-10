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

ls examples/$1/*/*/main.go | xargs -P $(($(nproc)*2)) -n 1 go build -o /dev/null
if [ $? -ne 0 ]; then
    echo -e "Failed to build examples"
    exit 1
fi
