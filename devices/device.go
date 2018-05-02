package devices

import (
	"github.com/vapor-ware/synse-sdk/sdk"
)

// AmtPower is the handler for the AMT controller.
var AmtPower = sdk.DeviceHandler{
	Type:  "power",
	Model: "AMT_POWER",

	Read:  amtPowerRead,
	Write: nil,
}

// amtPowerRead is the read handler function for amt power devices.
func amtPowerRead(device *sdk.Device) ([]*sdk.Reading, error) {

	readings := []*sdk.Reading{
		sdk.NewReading("state", "on"),
	}
	return readings, nil
}
