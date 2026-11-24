# sdk-e2e

End-to-end tests of the generated SDKs against a real Xen Orchestra, using the local build (never a published package).

Each language in `examples/` (go, nodejs, python, dotnet, and yaml as the SDK-free reference) runs the same scenario: read a pool, SR, network and template, create one VM, and export the **same output keys**. `expectations.json` holds the checks on those outputs. After `up`, `scripts/check.py` checks every language against them and compares the resource state between languages.

## Requirements

- `go`, `node`/`npm`, `uv`, `dotnet` with the ASP.NET Core runtime (`aspnet-runtime` on Arch), `pulumi`
- The XO env vars (`xoctx`): `XOA_URL`, `XOA_TOKEN`, `XOA_INSECURE`, `XOA_POOL`, `XOA_STORAGE`, `XOA_NETWORK`, `XOA_TEMPLATE`, optionally `XOA_TEST_PREFIX`. The provider reads them directly, no secret is written to disk.

## Usage

```sh
make provider build_sdks                  # from the repo root
sdk-e2e/scripts/run.sh all compile        # type-check every example, no XO needed
sdk-e2e/scripts/run.sh all preview        # + pulumi preview
sdk-e2e/scripts/run.sh all up             # + create one VM per language, export outputs, destroy
python3 sdk-e2e/scripts/check.py          # check the outputs of `up`
sdk-e2e/scripts/schema-diff.sh [base-ref] # schema changes since the last tag not covered yet
```

Replace `all` by `go`, `nodejs`, `python`, `dotnet` or `yaml` to run one language. `KEEP=1 run.sh <lang> up` skips the destroy. Logs and results land in `sdk-e2e/.out/<lang>/`, the Pulumi state in `sdk-e2e/.state/` (both ignored).

## Testing an upgrade with Claude Code

The `validate-sdks` skill (`.claude/skills/validate-sdks/`) does the whole loop: build, diff the schema since the last tag, add the new fields and functions to every example (one agent per language), run them, triage failures, and write `.out/report.md`. A change is tested once: its schema key goes to `covers` (or `skipped`, with a reason) in `expectations.json`, and later diffs ignore it.

Ready-to-use prompt:

```text
/validate-sdks up
Validate the generated SDKs for this provider upgrade. The `xoctx` env vars are already set.
Cover every schema change since the last tag (or mark it skipped with a reason), run all languages,
and give me the report: covered/skipped changes, language x step matrix, and each problem classified
(sdk-codegen, provider-mapping, upstream, test, infra) with the log line. Do not commit, do not post on the PR.
```

Use `compile` or `preview` instead of `up` for a run that creates no VM.
