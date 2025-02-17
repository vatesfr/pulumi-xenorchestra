# Xen Orchestra Provider

The Xen Orchestra Provider lets you manage [Xen Orchestra](https://github.com/vatesfr/xen-orchestra) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @vates/pulumi-xenorchestra
```

or `yarn`:

```bash
yarn add @vates/pulumi-xenorchestra
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-xenorchestra
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/vatesfr/pulumi-xenorchestra/sdk
```

<!-- ### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Foo
``` -->

## Configuration

Use `pulumi config set xenorchestra:<option>` or pass options to the constructor of new xenorchestra.Provider.

The following configuration points are available for the `xenorchestra` provider:

- `xenorchestra:url` (environment: `XOA_URL`) - the URL for the Xen Orchestra websockets endpoint. Starts with `wss://`
Set either:
- `xenorchestra:username` (environment: `XOA_USERNAME`) - the username for Xen Orchestra
- `xenorchestra:password` (environment: `XOA_PASSWORD`) - the password for Xen Orchestra
Or:
- `xenorchestra:token` (environment: `XOA_TOKEN`) - API token for Xen Orchestra

- `xenorchestra:insecure` (environment: `XOA_INSECURE`) - set to any value to disable SSL verification, false by default. Only use if you are using a self-signed certificate and know what you are doing.

## Usage

See [examples.md](./docs/examples.md) for examples.

## Building

See [contributing](CONTRIBUTING.md) for details on how to build and contribute to this provider.

## Reference
TBD
<!--
For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/foo/api-docs/). -->
