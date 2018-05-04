package devices

import (
	"fmt"
	"strings"

	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/logger"
)

// AmtBootTarget handler for setting an amt device's boot target.
var AmtBootTarget = sdk.DeviceHandler{
	Type:  "boot_target",
	Model: "amt-boot-target",

	Read:  bootTargetRead,
	Write: bootTargetWrite,
}

// bootTargetRead gets the current power state of the AMT device
func bootTargetRead(device *sdk.Device) ([]*sdk.Reading, error) {

	readings := []*sdk.Reading{
		sdk.NewReading("state", "default"),
	}
	return readings, nil
}

func bootTargetWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Raw
	// When writing to a boot_target device, we always expect there to be
	// raw data specified. If there isn't, we return an error.
	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'raw', but required")
	}

	if action == "state" {
		cmd := string(raw[0])

		switch strings.ToLower(cmd) {
		case "cd":
			logger.Debug("Set Boot Target to cd")
		case "hd":
			logger.Debug("Set Boot Target to hd")
		default:
			return fmt.Errorf("unsupported command for amt power 'state' action: %s", cmd)
		}

	} else {
		// If we reach here, then the specified action is not supported.
		return fmt.Errorf("action '%s' is not supported for bmc power devices", action)
	}

	return nil
}
