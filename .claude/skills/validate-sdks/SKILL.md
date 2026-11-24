---
name: validate-sdks
description: Validate the generated SDKs (go, nodejs, python, dotnet, plus yaml as reference) after a provider/upstream upgrade - extend the sdk-e2e examples with what changed in the schema since the last tag, run them against a real Xen Orchestra with the local build, and report failures and cross-SDK divergences. Use for "validate the SDKs", "test the generated SDKs", "e2e the upgrade", /validate-sdks [compile|preview|up] [base-ref].
---

# Validate the generated SDKs

Arguments: `[compile|preview|up]` (default `preview`; `up` creates one real VM per language, then destroys it) and an optional base ref for the schema diff (default: latest tag).

Everything lives in `sdk-e2e/`:
- `examples/<lang>/`: one program per language, all doing the same scenario and exporting the **same output keys** (camelCase in every language). They always use the local build: go `replace`, nodejs `file:` + `install-links`, python local path, dotnet `ProjectReference`, yaml uses the provider only.
- `expectations.json`: the contract. `outputs` = checks per output key; `covers` = schema keys an output validates; `skipped` = schema keys deliberately not tested, with the reason; `stateIgnore` = state paths allowed to differ between languages.
- `scripts/schema-diff.sh [base]`: schema changes since the base, minus `covers` and `skipped`. JSON on stdout.
- `scripts/run.sh <lang|all> <compile|preview|up>`: runs one or every example in parallel. Logs and results go to `sdk-e2e/.out/<lang>/`.
- `scripts/check.py`: checks `.out/*/outputs.json` against the expectations and compares the state between languages. Writes `.out/results.json`, exit 1 on failure.

Per-language conventions and pitfalls: [languages.md](languages.md). Read it before editing an example or briefing an agent.

## Steps

1. **Preconditions** (stop and tell the user if missing): `XOA_URL` and `XOA_TOKEN` set (the user runs `xoctx`; the provider reads `XOA_*` itself, never write the token to a file); `XOA_POOL`, `XOA_STORAGE`, `XOA_NETWORK`, `XOA_TEMPLATE` set; for dotnet, `dotnet --list-runtimes` shows `Microsoft.AspNetCore.App`.
2. **Build**: `make provider build_sdks` from the repo root. Do not use `make build` or `install_sdks`: they register a global yarn link and nuget source the examples do not need.
3. **Diff**: `sdk-e2e/scripts/schema-diff.sh [base] > sdk-e2e/.out/schema-diff.json`. If `changes` is empty, go to step 6.
4. **Spec** (you, not an agent: it needs the whole picture). For each change, decide:
   - testable → add or extend an output in `expectations.json` with its checks and `covers` (exact schema keys), and write what each program must do in `sdk-e2e/.out/spec.md`: which resource or invoke, which arguments, which output key, expected value. Prefer setting a new input and reading it back, and calling a new function with a filter that matches the test pool.
   - not testable here (needs hardware we lack, destructive, pure docs) → `skipped` with a one-line reason.
   - removed or changed type → make sure an output exercises it: these are the breaking changes users hit.
   Keep the scenario to one VM per language. New data sources are cheap (read-only); new resources must be destroyable.
   Edit the yaml example yourself: it has no SDK, so it is the reference that tells SDK bugs apart from provider bugs.
5. **Implement, in parallel**: launch 4 agents in one message (go, nodejs, python, dotnet) with the brief below, then wait for them.
6. **Run**: `sdk-e2e/scripts/run.sh all <mode>`, then `python3 -I sdk-e2e/scripts/check.py` (only for `up`: the other modes produce no outputs). In `preview` mode, read each `.out/<lang>/preview.log` instead: look for errors and output values whose type does not match the schema.
7. **Triage**: group failures by symptom (the same error in several languages = one problem). Triage a group yourself when it is obvious; otherwise launch one agent per group, in parallel, with the logs, `results.json`, and the classification below. Fix only test bugs (and rerun that language); never work around an SDK or provider bug in an example.
8. **Report**: write `sdk-e2e/.out/report.md` and give the user a short version: base ref, schema changes covered/skipped, a language × step matrix, each problem with its classification and evidence (log path, error line), and the suggested next action. Offer to post it on the PR; post only if the user says yes.

## Brief for an implementation agent (step 5)

> Update `sdk-e2e/examples/<lang>/` so it implements `sdk-e2e/.out/spec.md` and exports every output key of `sdk-e2e/expectations.json` with exactly that name. Read `.claude/skills/validate-sdks/languages.md` for the <lang> conventions and `sdk/<lang>/` for the generated API. Write the code a user would write: idiomatic, using the generated types. Iterate with `sdk-e2e/scripts/run.sh <lang> compile` until it passes. Do not touch other languages, `expectations.json`, the SDK, or the provider. If the idiomatic code cannot compile because of the generated SDK (missing type, wrong type, wrong name), do NOT work around it: keep the closest compiling code, and report it. Reply with: what you changed, the compile status, and any suspected SDK problem with the file and line in `sdk/<lang>/`.

## Classification (step 7)

- **sdk-codegen**: fails in one language, yaml passes → generated SDK (types, names, (de)serialisation). Point at `sdk/<lang>/…`.
- **provider-mapping**: wrong in every language including yaml, and the schema contradicts what the provider returns → `provider/resources.go` overrides or the bridge.
- **upstream**: the Terraform provider itself behaves that way (`upstream/xoa/…`).
- **test**: the example or the expectation is wrong → fix it and rerun.
- **infra**: XO, network, missing runtime, timeout → rerun, or tell the user what to install.

Known issue found by the first run (keep it until fixed): `Vm.memoryMax` is `number` in the schema (override in `provider/resources.go`, field `TypeInt` upstream), but the provider returns it as a string (`"1073741824"`). Go fails with `expected an float64, got a string`.
