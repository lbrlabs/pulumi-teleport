package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/pf/tfgen"
	teleport "github.com/lbrlabs/pulumi-teleport/provider"
)

func main() {
	tfgen.Main("teleport", teleport.Provider())
}
