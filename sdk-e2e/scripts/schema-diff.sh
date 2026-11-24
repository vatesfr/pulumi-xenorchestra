#!/usr/bin/env bash
# Lists schema changes between a base ref (default: latest tag) and the working tree,
# minus what sdk-e2e/expectations.json already covers (outputs[].covers) or skips on purpose (skipped).
# Usage: schema-diff.sh [base-ref] [--all]   (--all: do not filter covered entries)
# Output: JSON array of {key, change: added|removed|changed, old, new} on stdout.
set -euo pipefail
cd "$(git rev-parse --show-toplevel)"
SCHEMA=provider/cmd/pulumi-resource-xenorchestra/schema.json
BASE=${1:-$(git describe --tags --abbrev=0)}
FILTER=1; [[ ${2:-} == --all || ${1:-} == --all ]] && FILTER=0
[[ $BASE == --all ]] && BASE=$(git describe --tags --abbrev=0)

# Flatten the schema to {"<token>": sig, "<token>#<section>.<prop>": sig}
FLATTEN='
def sig($k; $req): {type: (.type // ."$ref"), items: (.items.type // .items."$ref"),
                required: ($req | index([$k]) != null) , deprecated: (.deprecationMessage != null),
                secret: (.secret // false)};
def props($tok; $sec; $obj):
  ($obj.properties // {}) as $p | ($obj.required // []) as $r
  | $p | to_entries[] | .key as $k | {key: "\($tok)#\($sec).\($k)", value: (.value | sig($k; $r))};
[ (.resources // {} | to_entries[] | .key as $t
    | {key: $t, value: "resource"},
      props($t; "inputs"; {properties: .value.inputProperties, required: .value.requiredInputs}),
      props($t; "outputs"; {properties: .value.properties, required: .value.required})),
  (.functions // {} | to_entries[] | .key as $t
    | {key: $t, value: "function"},
      props($t; "inputs"; (.value.inputs // {})),
      props($t; "outputs"; (.value.outputs // {}))),
  (.types // {} | to_entries[] | .key as $t
    | {key: $t, value: "type"}, props($t; "properties"; .value)),
  props("provider"; "inputs"; {properties: .provider.inputProperties, required: .provider.requiredInputs})
] | from_entries'

old=$(git show "$BASE:$SCHEMA" | jq "$FLATTEN")
new=$(jq "$FLATTEN" "$SCHEMA")
covered='[]'
if [[ $FILTER == 1 && -f sdk-e2e/expectations.json ]]; then
  covered=$(jq '[.outputs[].covers[]?] + (.skipped // {} | keys)' sdk-e2e/expectations.json)
fi

jq -n --argjson old "$old" --argjson new "$new" --argjson covered "$covered" --arg base "$BASE" '
  [ ($new | keys[]) as $k | select($old[$k] == null) | {key: $k, change: "added", new: $new[$k]} ]
  + [ ($old | keys[]) as $k | select($new[$k] == null) | {key: $k, change: "removed", old: $old[$k]} ]
  + [ ($new | keys[]) as $k | select($old[$k] != null and $old[$k] != $new[$k])
    | {key: $k, change: "changed", old: $old[$k], new: $new[$k]} ]
  # covers entries match exact keys only: covering a function does not cover its new fields
  | map(select(.key as $k | $covered | index([$k]) | not))
  | {base: $base, changes: .}'
