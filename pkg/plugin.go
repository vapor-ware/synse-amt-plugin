package pkg

import (
	"log"

	"github.com/vapor-ware/synse-amt-plugin/pkg/devices"
	"github.com/vapor-ware/synse-amt-plugin/pkg/outputs"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// MakePlugin creates a new instance of the AMT plugin.
func MakePlugin() *sdk.Plugin {
	plugin := sdk.NewPlugin(
		sdk.CustomDeviceIdentifier(deviceIdentifier),
	)

	err := plugin.RegisterOutputTypes(
		&outputs.BootTarget,
		&outputs.PowerState,
	)
	if err != nil {
		log.Fatal(err)
	}

	plugin.RegisterDeviceHandlers(
		&devices.AmtBootTarget,
		&devices.AmtPower,
	)

	return plugin
}
