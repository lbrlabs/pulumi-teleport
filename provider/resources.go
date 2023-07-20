package teleport

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/lbrlabs/pulumi-teleport/provider/pkg/version"
	tfpfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/gravitational/teleport-plugins/terraform/shim"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	teleportPkg = "teleport"
	// modules:
	teleportMod = "index" // the teleport module
)

// teleportMember manufactures a type token for the teleport package and the given module and type.
func teleportMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(teleportPkg + ":" + mod + ":" + mem)
}

// teleportType manufactures a type token for the teleport package and the given module and type.
func teleportType(mod string, typ string) tokens.Type {
	return tokens.Type(teleportMember(mod, typ))
}

// teleportDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the teleport package and names the file by simply lower casing the data
// source's first character.
func teleportDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return teleportMember(mod+"/"+fn, res)
}

// teleportResource manufactures a standard resource token given a module and resource name.
// It automatically uses the teleport package and names the file by simply lower casing the resource's
// first character.
func teleportResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return teleportType(mod+"/"+fn, res)
}

//go:embed cmd/pulumi-resource-teleport/bridge-metadata.json
var bridgeMetadata []byte

func boolRef(b bool) *bool {
	return &b
}

// Provider returns additional overlaid schema and metadata associated with the tls package.
func Provider() tfpfbridge.ProviderInfo {
	info := tfbridge.ProviderInfo{
		Name:              "teleport",
		DisplayName:       "Teleport",
		Publisher:         "lbrlabs",
		LogoURL:           "https://raw.githubusercontent.com/lbrlabs/pulumi-teleport/main/assets/teleport.png", // nolint[:lll]
		PluginDownloadURL: "github://api.github.com/lbrlabs",
		Description:       "A Pulumi package to create Teleport resources in Pulumi programs.",
		Keywords:          []string{"pulumi", "teleport", "lbrlabs"},
		License:           "Apache-2.0",
		Homepage:          "https://pulumi.io",
		Repository:        "https://github.com/pulumi/pulumi-tls",
		Version:           version.Version,
		GitHubOrg:         "gravitational",
		Resources:         map[string]*tfbridge.ResourceInfo{},
		DataSources:       map[string]*tfbridge.DataSourceInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@lbrlabs/pulumi-teleport",
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
			PackageName: "lbrlabs_pulumi_teleport",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/lbrlabs/pulumi-%[1]s/sdk/", teleportPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				teleportPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Lbrlabs.PulumiPackage",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	return tfpfbridge.ProviderInfo{
		ProviderInfo: info,
		NewProvider:  shim.NewProvider,
	}
}
