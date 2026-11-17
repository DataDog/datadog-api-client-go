#!/usr/bin/env bash
# Tidy every split module under packages/; the generator writes no go.sum, so this writes it.
set -euo pipefail
shopt -s nullglob

for dir in packages/*/; do
    (cd "$dir" && GOWORK=off go mod tidy)
done
