#!/usr/bin/env bash
# Runs one SDK example (or all of them in parallel) against the locally built provider and SDKs.
# Usage: run.sh <go|nodejs|python|dotnet|yaml|all> [compile|preview|up]   (default: preview)
#   compile  install deps + type-check, no XO needed
#   preview  compile + pulumi preview (reads XO through the invokes)
#   up       compile + pulumi up, save outputs/state, then always destroy (KEEP=1 skips the destroy)
# Needs: `make provider build_sdks` done, and XOA_* env vars (xoctx) for preview/up.
# Writes: sdk-e2e/.out/<lang>/{status,<step>.log,outputs.json,state.json}
set -uo pipefail
E2E=$(cd "$(dirname "$0")/.." && pwd)
REPO=$(dirname "$E2E")
LANGS=(go nodejs python dotnet yaml)
lang=${1:?usage: run.sh <lang|all> [compile|preview|up]}
mode=${2:-preview}
[[ $mode =~ ^(compile|preview|up)$ ]] || { echo "unknown mode: $mode" >&2; exit 2; }

if [[ $lang == all ]]; then
  pids=()
  for l in "${LANGS[@]}"; do "$0" "$l" "$mode" & pids+=($!); done
  rc=0; for p in "${pids[@]}"; do wait "$p" || rc=1; done
  exit $rc
fi
[[ " ${LANGS[*]} " == *" $lang "* ]] || { echo "unknown lang: $lang" >&2; exit 2; }
[[ -x $REPO/bin/pulumi-resource-xenorchestra ]] || { echo "provider not built: run 'make provider build_sdks'" >&2; exit 2; }
if [[ $mode != compile ]]; then
  : "${XOA_URL:?XOA_URL not set, run xoctx first}" "${XOA_TOKEN:?XOA_TOKEN not set, run xoctx first}"
fi

export PATH="$REPO/bin:$PATH"   # pulumi picks the locally built provider as an ambient plugin
export PULUMI_BACKEND_URL="file://$E2E/.state" PULUMI_CONFIG_PASSPHRASE="" PULUMI_SKIP_UPDATE_CHECK=true
DIR=$E2E/examples/$lang
OUT=$E2E/.out/$lang
mkdir -p "$E2E/.state"; rm -rf "$OUT"; mkdir -p "$OUT"
cd "$DIR"

step() { # step <name> <cmd...>: logs to $OUT/<name>.log, records ok/FAIL in $OUT/status
  local name=$1; shift
  echo "[$lang] $name..."
  if "$@" >"$OUT/$name.log" 2>&1; then echo "$name ok" >>"$OUT/status"
  else echo "$name FAIL" >>"$OUT/status"; echo "[$lang] $name FAILED: $OUT/$name.log"; return 1; fi
}

compile() {
  case $lang in
    go) go mod tidy && go vet . ;;
    # drop the copied SDK so npm re-copies the freshly built one (same version string every build)
    nodejs) rm -rf node_modules/@vates/pulumi-xenorchestra && npm install --no-audit --no-fund && npx tsc --noEmit ;;
    # --seed: pulumi calls `venv/bin/python -m pip list`, a bare uv venv has no pip
    python) { [[ -x venv/bin/pip ]] || uv venv -q --seed --clear venv; } \
      && uv pip install -q -p venv/bin/python -r requirements.txt --reinstall-package pulumi_xenorchestra \
      && venv/bin/python -m py_compile __main__.py && venv/bin/python -c 'import pulumi_xenorchestra' ;;
    dotnet) dotnet build -nologo -v q ;;
    yaml) true ;; # no SDK: validated by preview
  esac
}

stack() {
  pulumi stack select --create e2e --non-interactive || return 1
  if [[ $lang == yaml ]]; then
    pulumi config set poolName "$XOA_POOL" && pulumi config set storageName "$XOA_STORAGE" \
      && pulumi config set networkName "$XOA_NETWORK" && pulumi config set templateName "$XOA_TEMPLATE" \
      && pulumi config set prefix "${XOA_TEST_PREFIX:-}"
  fi
}

up() {
  pulumi up --yes --non-interactive --skip-preview --diff || return 1
  pulumi stack output --json --show-secrets >"$OUT/outputs.json" && pulumi stack export >"$OUT/state.json"
}

step compile compile || exit 1
[[ $mode == compile ]] && exit 0
step stack stack || exit 1
if [[ $mode == preview ]]; then step preview pulumi preview --non-interactive --diff; exit; fi
step up up; rc=$?
[[ ${KEEP:-0} == 1 ]] || step destroy pulumi destroy --yes --non-interactive || rc=1
exit $rc
