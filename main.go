package main

import (
	"github.com/Triple-Whale/terraform-provider-jetstream/jetstream"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: jetstream.Provider})
}
