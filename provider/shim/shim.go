package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/gravitational/teleport-plugins/terraform/provider"
)

func NewProvider() tfpf.Provider {
	return provider.New()
}