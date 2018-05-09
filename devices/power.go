package devices

import (
	"os/exec"

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

	cmd := exec.Command("python", "scripts/power.py", device.Data["ip"], // nolint: gas
		device.Data["password"], "status")

	out, err := cmd.Output()

	if err != nil {
		logger.Errorf("Error: %s\n", cmd.Stderr)
		return nil, err
	}

	readings := []*sdk.Reading{
		sdk.NewReading("state", string(out)),
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
		logger.Error("no values specified for 'raw', but required")
		return nil
	}

	if action == "state" {
		commandName := string(raw[0])
		cmd := exec.Command("python", "scripts/power.py", device.Data["ip"], // nolint: gas
			device.Data["password"], commandName)

		_, err := cmd.Output()

		if err != nil {
			logger.Errorf("Error: %s\n", cmd.Stderr)
			return err
		}

	} else {
		// If we reach here, then the specified action is not supported.
		logger.Error("action '%s' is not supported for AMT power devices", action)
		return nil
	}

	return nil
}
