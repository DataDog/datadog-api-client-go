#!/usr/bin/env bash
# Point tests/go.mod at every split module under packages/ and tidy it; the generator doesn't write it.
set -euo pipefail
shopt -s nullglob

[ -f tests/scenarios/package_mappings.go ] || exit 0

edits=()
modules=" "
for gomod in packages/*/go.mod; do
    dir=$(dirname "$gomod")
    module=$(awk '$1 == "module" { print $2; exit }' "$gomod")
    version=$(sed -n 's/^const Version = "\(.*\)"$/\1/p' "$dir/version.go")
    edits+=(-require="$module@v$version" -replace="$module=../$dir")
    modules+="$module "
done

for module in $(awk '$1 == "replace" && $2 ~ "^github.com/DataDog/datadog-api-client-go/packages/" { print $2 }' tests/go.mod); do
    [[ "$modules" == *" $module "* ]] || edits+=(-droprequire="$module" -dropreplace="$module")
done

cd tests
go mod edit "${edits[@]}"
GOWORK=off go mod tidy
