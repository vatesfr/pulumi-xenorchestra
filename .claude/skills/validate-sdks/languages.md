# Per-language notes for sdk-e2e/examples

Output keys are camelCase in every language (`vmCpus`), whatever the language's own naming.

| | go | nodejs | python | dotnet | yaml |
|---|---|---|---|---|---|
| Program | `main.go` | `index.ts` | `__main__.py` | `Program.cs` | `Pulumi.yaml` |
| Local SDK via | `replace …/sdk/v2 => ../../../sdk` in `go.mod` | `file:../../../sdk/nodejs/bin` + `.npmrc install-links=true` | `../../../sdk/python/bin` in `requirements.txt` | `ProjectReference` to `sdk/dotnet/Pulumi.Xenorchestra.csproj` | provider only (`bin/` on PATH) |
| SDK sources to read | `sdk/go/xenorchestra/` | `sdk/nodejs/` (`types/input.ts`, `types/output.ts`) | `sdk/python/pulumi_xenorchestra/` | `sdk/dotnet/` (`Inputs/`, `Outputs/`) | `provider/cmd/pulumi-resource-xenorchestra/schema.json` |
| Names | `NameLabel`, `xoa.VmArgs`, `xoa.VmDiskArgs` | `nameLabel`, plain objects | `name_label`, `xo.VmDiskArgs` | `NameLabel`, `new VmArgs { … }` | schema names (`nameLabel`) |
| Invoke | `xoa.GetXoaPool(ctx, &args)` (plain), `LookupXoaNetwork` for data sources named like a resource | `getXoaPoolOutput({...})` | `xo.get_xoa_pool(...)` (plain) | `GetXoaPool.Invoke(new() {...})` → `Output<T>` | `fn::invoke` |
| Export | `ctx.Export("key", v)` | `export const key = v` | `pulumi.export("key", v)` | returned `Dictionary` | `outputs:` |
| Compile check | `go vet` | `tsc --noEmit` | import + `py_compile` (no type checker: check names by reading the SDK) | `dotnet build` | none, preview validates |

Pitfalls:
- Go: optional plain invoke args are pointers (`&pool.Id`, `pulumi.StringRef`), required ones are values; the compiler tells which. A typed output that the provider returns with another type fails at runtime only (`expected an float64, got a string`), not in `go vet`.
- nodejs: `run.sh` deletes and re-copies `node_modules/@vates/pulumi-xenorchestra` every run, because the dev version string never changes. Never use `yarn link`/symlinks: a second `@pulumi/pulumi` copy breaks resources.
- python: the venv is created by `uv venv --seed` (pulumi needs pip in it). Mistyped keyword arguments only fail at runtime: check the `__init__` signature in the SDK.
- dotnet: needs the ASP.NET Core runtime (`aspnet-runtime` on Arch). `RollForward=Major` lets the net8.0 target run on a newer runtime.
- yaml: cannot read env vars: `run.sh` copies `XOA_POOL`… into the stack config (`poolName`, `storageName`, `networkName`, `templateName`, `prefix`). A config key must not share a name with a variable.
- All: the disk must be at least the template disk size (8 GiB for `Alpine 3.23`), XO refuses to shrink a VDI.
- All: VM names are `${XOA_TEST_PREFIX}pulumi-e2e-<lang>` so the 5 runs never collide.
