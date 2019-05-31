package pkg

import (
	log "github.com/sirupsen/logrus"
	"github.com/vapor-ware/synse-amt-plugin/pkg/devices"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// MakePlugin creates a new instance of the AMT plugin.
func MakePlugin() *sdk.Plugin {
	plugin, err := sdk.NewPlugin(
		sdk.CustomDeviceIdentifier(deviceIdentifier),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register device handlers.
	err = plugin.RegisterDeviceHandlers(
		&devices.AmtBootTarget,
		&devices.AmtPower,
	)
	if err != nil {
		log.Fatal(err)
	}

	return plugin
}
