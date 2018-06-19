package devices

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/scripts"
)

// AmtPower is the handler for the AMT power controller.
var AmtPower = sdk.DeviceHandler{
	Name:  "power",
	Read:  amtPowerRead,
	Write: amtPowerWrite,
}

// amtPowerRead gets the current power state of the AMT device
func amtPowerRead(device *sdk.Device) ([]*sdk.Reading, error) {

	ip := fmt.Sprint(device.Data["ip"])
	pass := fmt.Sprint(device.Data["password"])

	cmd := scripts.NewCommand("python", "scripts/power.py", ip, pass, "status")
	err := cmd.Run()
	if err != nil {
		log.Errorf("error: %s", cmd.Stderr())
		return nil, err
	}

	readings := []*sdk.Reading{
		device.GetOutput("power.state").MakeReading(cmd.Stdout()),
	}
	return readings, nil
}

// amtPowerWrite sets the power state of the AMT device
func amtPowerWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Data

	// When writing to a AMT Power device, we always expect there to be
	// raw data specified. If there isn't, we return an error.
	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'raw', but required")
	}

	if action == "state" {
		commandName := string(raw)

		ip := fmt.Sprint(device.Data["ip"])
		pass := fmt.Sprint(device.Data["password"])

		cmd := scripts.NewCommand("python", "scripts/power.py", ip, pass, commandName)
		err := cmd.Run()
		if err != nil {
			log.Errorf("error: %s", cmd.Stderr())
			return err
		}
		return nil

	}

	// If we reach here, then the specified action is not supported.
	return fmt.Errorf("action '%s' is not supported for AMT power devices", action)
}
