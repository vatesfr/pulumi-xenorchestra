// Copyright 2016-2023, Pulumi Corporation.
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
	"fmt"
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	xenorchestra "github.com/vatesfr/terraform-provider-xenorchestra/xoa"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

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

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(resource.PropertyMap, shim.ResourceConfig) error {
	return nil
}

//go:embed cmd/pulumi-resource-xenorchestra/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		// Instantiate the Terraform provider
		P:    shimv2.NewProvider(xenorchestra.Provider()),
		Name: "xenorchestra",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Vates",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://github.com/vatesfr/xen-orchestra/raw/587da7b1336da1acbd7783555b863a9fd00e2893/@xen-orchestra/web/public/favicon.svg", //nolint:all
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/vatesfr/pulumi-xenorchestra",
		Description:       "A Pulumi package for creating and managing Xen Orchestra cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "xenorchestra", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/vatesfr/pulumi-xenorchestra",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg:    "vatesfr",
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		Config:       map[string]*tfbridge.SchemaInfo{
			// "url": {
			//     Default: &tfbridge.DefaultInfo{
			//         EnvVars: []string{"XOA_URL"},
			//     }
			// },
			// "username": {
			//     Default: &tfbridge.DefaultInfo{
			//         EnvVars: []string{"XOA_USERNAME"},
			//     }
			// },
			// "password": {
			//     Default: &tfbridge.DefaultInfo{
			//         EnvVars: []string{"XOA_PASSWORD"},
			//     }
			// },
			// "token": {
			//     Default: &tfbridge.DefaultInfo{
			//         EnvVars: []string{"XOA_TOKEN"},
			//     }
			// },
			// "insecure": {
			//     Default: &tfbridge.DefaultInfo{
			//         EnvVars: []string{"XOA_INSECURE"},
			//     }
			// }
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"xenorchestra_acl":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Acl")},
			"xenorchestra_bonded_network": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "XoaBondedNetwork")},
			"xenorchestra_cloud_config":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudConfig")},
			"xenorchestra_network":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "XoaNetwork")},
			"xenorchestra_resource_set":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ResourceSet")},
			"xenorchestra_vdi":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Vdi")},
			"xenorchestra_vm":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Vm")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"xenorchestra_cloud_config": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaCloudConfig")},
			"xenorchestra_host":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaHost")},
			"xenorchestra_hosts":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaHosts")},
			"xenorchestra_network":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaNetwork")},
			"xenorchestra_pif":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaPif")},
			"xenorchestra_pool":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaPool")},
			"xenorchestra_resource_set": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaResourceSet")},
			"xenorchestra_sr":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaStorageRepository")},
			"xenorchestra_template":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaTemplate")},
			"xenorchestra_user":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaUser")},
			"xenorchestra_vdi":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaVdi")},
			"xenorchestra_vms":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getXoaVms")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@vates/pulumi-xenorchestra",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/vatesfr/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource
	// tokens, and apply auto aliasing for full backwards compatibility.  For more
	// information, please reference:
	// https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("xenorchestra_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
