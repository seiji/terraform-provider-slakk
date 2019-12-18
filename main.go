package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/seiji/terraform-provider-slakk/slack"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: slack.Provider})
}
