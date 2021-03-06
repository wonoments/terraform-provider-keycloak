package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/tazjin/terraform-provider-keycloak/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
