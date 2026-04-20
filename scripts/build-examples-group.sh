#!/usr/bin/env bash
set -e

GROUP=${1:?usage: build-examples-group.sh <group-number>}
EXAMPLES_PER_GROUP=${EXAMPLES_PER_GROUP:-400}

ALL_EXAMPLES=()
while IFS= read -r f; do ALL_EXAMPLES+=("$f"); done < <(find examples -mindepth 3 -maxdepth 3 -name "*.go" | sort)
TOTAL=${#ALL_EXAMPLES[@]}

START=$(( (GROUP - 1) * EXAMPLES_PER_GROUP ))
SLICE=("${ALL_EXAMPLES[@]:$START:$EXAMPLES_PER_GROUP}")

echo "Building group $GROUP: examples $((START + 1))-$((START + ${#SLICE[@]})) of $TOTAL"

for f in "${SLICE[@]}"; do
    df="build/$(dirname "$f")/$(basename "$f" .go)"
    mkdir -p "$df"
    cp "$f" "$df/main.go"
done

if go build -o /dev/null -ldflags "-s -w -linkmode internal" ./build/examples/*/*/*; then
    echo "Examples are buildable"
else
    echo "Failed to build examples"
    exit 1
fi
