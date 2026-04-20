#!/usr/bin/env bash
set -e

EXAMPLES_PER_GROUP=${EXAMPLES_PER_GROUP:-400}

ALL_EXAMPLES=()
while IFS= read -r f; do ALL_EXAMPLES+=("$f"); done < <(find examples -mindepth 3 -maxdepth 3 -name "*.go" | sort)
TOTAL=${#ALL_EXAMPLES[@]}
NUM_GROUPS=$(( (TOTAL + EXAMPLES_PER_GROUP - 1) / EXAMPLES_PER_GROUP ))

echo "Total examples: $TOTAL, groups: $NUM_GROUPS (max $EXAMPLES_PER_GROUP per group)"

MATRIX_JSON="["
for i in $(seq 1 "$NUM_GROUPS"); do
    [ "$i" -gt 1 ] && MATRIX_JSON+=","
    MATRIX_JSON+="$i"
done
MATRIX_JSON+="]"

echo "matrix=$MATRIX_JSON"
if [ -n "${GITHUB_OUTPUT:-}" ]; then
    echo "matrix=$MATRIX_JSON" >> "$GITHUB_OUTPUT"
fi
