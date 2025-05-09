// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xenorchestra

import (
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	xenorchestra "github.com/vatesfr/terraform-provider-xenorchestra/xoa" // Import the upstream provider

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/vatesfr/pulumi-xenorchestra/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "xenorchestra"
	// modules:
	mainMod = "index" // the xenorchestra module
)

// @gCyrille should we keep this function?
// func preConfigureCallback(resource.PropertyMap, shim.ResourceConfig) error {
// 	return nil
// }

//go:embed cmd/pulumi-resource-xenorchestra/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		// Instantiate the Terraform provider
		//
		// The [pulumi-terraform-bridge](https://github.com/pulumi/pulumi-terraform-bridge) supports 3
		// types of Terraform providers:
		//
		// 1. Providers written with the terraform-plugin-sdk/v1:
		//
		//    If the provider you are bridging is written with the terraform-plugin-sdk/v1, then you
		//    will need to adapt the boilerplate:
		//
		//    - Change the import "shimv2" to "shimv1" and change the associated import to
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1".
		//
		//    You can then proceed as normal.
		//
		// 2. Providers written with terraform-plugin-sdk/v2:
		//
		//    This boilerplate is already geared towards providers written with the
		//    terraform-plugin-sdk/v2, since it is the most common provider framework used. No
		//    adaptions are needed.
		//
		// 3. Providers written with terraform-plugin-framework:
		//
		//    If the provider you are bridging is written with the terraform-plugin-framework, then
		//    you will need to adapt the boilerplate:
		//
		//    - Remove the `shimv2` import and add:
		//
		//      	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
		//
		//    - Replace `shimv2.NewProvider` with `pfbridge.ShimProvider`.
		//
		//    - In provider/cmd/pulumi-tfgen-xenorchestra/main.go, replace the
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen" import with
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen". Remove the `version.Version`
		//      argument to `tfgen.Main`.
		//
		//    - In provider/cmd/pulumi-resource-xenorchestra/main.go, replace the
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge" import with
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge". Replace the arguments to the
		//      `tfbridge.Main` so it looks like this:
		//
		//      	tfbridge.Main(context.Background(), "xenorchestra", xenorchestra.Provider(),
		//			tfbridge.ProviderMetadata{PulumiSchema: pulumiSchema})
		//
		//   Detailed instructions can be found at
		//   https://pulumi-developer-docs.readthedocs.io/projects/pulumi-terraform-bridge/en/latest/docs/guides/new-pf-provider.html
		//   After that, you can proceed as normal.
		//
		// This is where you give the bridge a handle to the upstream terraform provider. SDKv2
		// convention is to have a function at "github.com/vatesfr/terraform-provider-xenorchestra/xoa".New
		// which takes a version and produces a factory function. The provider you are bridging may
		// not do that. You will need to find the function (generally called in upstream's main.go)
		// that produces a:
		//
		// - *"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema".Provider (for SDKv2)
		// - *"github.com/hashicorp/terraform-plugin-sdk/v1/helper/schema".Provider (for SDKv1)
		// - "github.com/hashicorp/terraform-plugin-framework/provider".Provider (for plugin-framework)
		//
		//nolint:lll
		P: shimv2.NewProvider(xenorchestra.Provider()),

		Name:    "xenorchestra",
		Version: version.Version,
		// DisplayName is a way to be able to change the casing of the provider name when being
		// displayed on the Pulumi registry
		DisplayName: "",
		// Change this to your personal name (or a company name) that you would like to be shown in
		// the Pulumi Registry if this package is published there.
		Publisher: "Vates",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an PNG logo (100x100) for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://raw.githubusercontent.com/vatesfr/pulumi-xenorchestra/8c71624229d953d4fb7d4843d5483e53e21b9459/logo_xo.png", //nolint:all
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g. https://github.com/org/pulumi-provider-name/releases/download/v${VERSION}/
		PluginDownloadURL: "github://api.github.com/vatesfr/pulumi-xenorchestra",
		Description:       "A Pulumi package for creating and managing Xen Orchestra cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "xenorchestra", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/vatesfr/pulumi-xenorchestra",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this should
		// match the TF provider module's require directive, not any replace directives.
		GitHubOrg:    "vatesfr",
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		Config:       map[string]*tfbridge.SchemaInfo{},
		// If extra types are needed for configuration, they can be added here.

		ExtraTypes: map[string]schema.ComplexTypeSpec{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"xenorchestra_acl":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Acl")},
			"xenorchestra_bonded_network": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "XoaBondedNetwork")},
			"xenorchestra_cloud_config":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudConfig")},
			"xenorchestra_network":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "XoaNetwork")},
			"xenorchestra_resource_set":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ResourceSet")},
			"xenorchestra_vdi":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Vdi")},
			"xenorchestra_vm": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Vm"),
				Fields: map[string]*tfbridge.SchemaInfo{
					// memory_max and disk size should be a u64 since it describes bytes sizes.
					//
					// The Pulumi schema doesn't support u64 (or i64), so we
					// are converting to a float64, which should get us at
					// least 2^52 bits of precision while minimizing the
					// breaking change (int -> float) for our users.
					"memory_max": {
						Type: "number",
					},
					"disk": {Elem: &tfbridge.SchemaInfo{Fields: map[string]*tfbridge.SchemaInfo{
						"size": {
							Type: "number",
						},
					}}},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"xenorchestra_cloud_config": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaCloudConfig")},
			"xenorchestra_host": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaHost"),
				Fields: map[string]*tfbridge.SchemaInfo{
					// memory and memory_usage should be a u64 since it describes bytes sizes.
					//
					// The Pulumi schema doesn't support u64 (or i64), so we
					// are converting to a float64, which should get us at
					// least 2^52 bits of precision while minimizing the
					// breaking change (int -> float) for our users.
					"memory": {
						Type: "number",
					},
					"memory_usage": {
						Type: "number",
					},
				},
			},
			"xenorchestra_hosts":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaHosts")},
			"xenorchestra_network": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaNetwork")},
			"xenorchestra_pif":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaPif")},
			"xenorchestra_pool":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaPool")},
			"xenorchestra_resource_set": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaResourceSet"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"limit": {Elem: &tfbridge.SchemaInfo{Fields: map[string]*tfbridge.SchemaInfo{
						// quantity should be a u64 since it describes bytes sizes.
						//
						// The Pulumi schema doesn't support u64 (or i64), so we
						// are converting to a float64, which should get us at
						// least 2^52 bits of precision while minimizing the
						// breaking change (int -> float) for our users.
						"quantity": {
							Type: "number",
						},
					}}},
				},
			},
			"xenorchestra_sr": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaStorageRepository"),
				Fields: map[string]*tfbridge.SchemaInfo{
					// size, physical_usage and usage should be a u64 since it describes bytes sizes.
					//
					// The Pulumi schema doesn't support u64 (or i64), so we
					// are converting to a float64, which should get us at
					// least 2^52 bits of precision while minimizing the
					// breaking change (int -> float) for our users.
					"size": {
						Type: "number",
					},
					"physical_usage": {
						Type: "number",
					},
					"usage": {
						Type: "number",
					},
				},
			},
			"xenorchestra_template": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaTemplate")},
			"xenorchestra_user":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaUser")},
			"xenorchestra_vdi":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaVdi")},
			"xenorchestra_vms":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaVms")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@vates/pulumi-xenorchestra",
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
			// Enable modern PyProject support in the generated Python SDK.
			PyProject: struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			// Set where the SDK is going to be published to.
			ImportBasePath: path.Join(
				"github.com/vatesfr/pulumi-xenorchestra/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			// Opt in to all available code generation features.
			GenerateResourceContainerTypes: true,
			GenerateExtraInputTypes:        true,
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
			// Use a wildcard import so NuGet will prefer the latest possible version.
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// MustComputeTokens maps all resources and datasources from the upstream provider into Pulumi.
	//
	// tokens.SingleModule puts every upstream item into your provider's main module.
	//
	// You shouldn't need to override anything, but if you do, use the [tfbridge.ProviderInfo.Resources]
	// and [tfbridge.ProviderInfo.DataSources].
	prov.MustComputeTokens(tokens.SingleModule("xenorchestra_", mainMod,
		tokens.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
