package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/vapor-ware/synse-amt-plugin/pkg"
	"github.com/vapor-ware/synse-sdk/sdk"
)

const (
	pluginName       = "intel amt"
	pluginMaintainer = "vaporio"
	pluginDesc       = "A simple plugin for communicating with Intel AMT enabled machines"
	pluginVcs        = "https://github.com/vapor-ware/synse-amt-plugin"
)

func main() {
	// Set the plugin metadata
	sdk.SetPluginInfo(
		pluginName,
		pluginMaintainer,
		pluginDesc,
		pluginVcs,
	)

	plugin := pkg.MakePlugin()

	// Run the plugin
	if err := plugin.Run(); err != nil {
		log.Fatal(err)
	}
}
