package devices

import (
	"fmt"
	"strings"

	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/logger"
)

// AmtPower is the handler for the AMT power controller.
var AmtPower = sdk.DeviceHandler{
	Type:  "power",
	Model: "amt-power",

	Read:  amtPowerRead,
	Write: amtPowerWrite,
}

// amtPowerRead gets the current power state of the AMT device
func amtPowerRead(device *sdk.Device) ([]*sdk.Reading, error) {

	readings := []*sdk.Reading{
		sdk.NewReading("state", "on"),
	}
	return readings, nil
}

// amtPowerWrite sets the power state of the AMT device
func amtPowerWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Raw
	// When writing to a AMT Power device, we always expect there to be
	// raw data specified. If there isn't, we return an error.
	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'raw', but required")
	}

	if action == "state" {
		cmd := string(raw[0])

		switch strings.ToLower(cmd) {
		case "on":
			logger.Debug("AMT Power On")
		case "off":
			logger.Debug("AMT Power Off")
		case "reset":
			logger.Debug("AMT Power Resetting")
		case "cycle":
			logger.Debug("AMT Power Cycling")
		default:
			return fmt.Errorf("unsupported command for amt power 'state' action: %s", cmd)
		}

	} else {
		// If we reach here, then the specified action is not supported.
		return fmt.Errorf("action '%s' is not supported for AMT power devices", action)
	}

	return nil
}
