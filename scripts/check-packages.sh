#!/usr/bin/env bash
# Check that every split package under packages/ is its own module, and that no root module can contain them.
set -euo pipefail
shopt -s nullglob

status=0
if [ -f go.mod ]; then
    echo "go.mod at the repository root would make packages/ part of a root module"
    status=1
fi

for dir in packages/*/; do
    name=$(basename "$dir")
    if [ ! -f "$dir/go.mod" ]; then
        echo "packages/$name has no go.mod"
        status=1
        continue
    fi
    module=$(awk '$1 == "module" { print $2; exit }' "$dir/go.mod")
    if ! [[ "$module" =~ ^github\.com/DataDog/datadog-api-client-go/packages/${name}(/v([2-9]|[1-9][0-9]+))?$ ]]; then
        echo "packages/$name/go.mod declares module '$module'"
        status=1
    fi
done

if [ -d packages ]; then
    while IFS= read -r gomod; do
        echo "$gomod is not at packages/<name>/go.mod"
        status=1
    done < <(find packages -name go.mod | grep -Ev '^packages/[^/]+/go\.mod$' || true)
fi

exit $status
