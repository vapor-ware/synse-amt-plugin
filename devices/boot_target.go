package devices

import (
	"fmt"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/logger"
	"os/exec"
)

// AmtBootTarget is the handler for setting an amt device's boot target
var AmtBootTarget = sdk.DeviceHandler{
	Type:  "boot_target",
	Model: "amt-boot-target",

	Read:  nil,
	Write: bootTargetWrite,
}

// bootTargetWrite sets the amt boot target
func bootTargetWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Raw
	// When writing to a boot_target device, we always expect there to be
	// raw data specified. If there isn't, we return an error.
	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'raw', but required")
	}

	if action == "target" {
		target := string(raw[0])

		supportedTargets := map[string]bool {
			"pxe": true,
			"hd": true,
			"cd": true,
		}

		if supportedTargets[target] {
			logger.Info("setting boot target to ", target)
			cmd := exec.Command("python", "scripts/boot_target.py", device.Data["ip"], // nolint: gas
				device.Data["password"], target)

			_, err := cmd.Output()

			if err != nil {
				logger.Errorf("error: %s", cmd.Stderr)
				return err
			}

			return nil

		}

		// If we reach here, the specified boot target is not supported
		return fmt.Errorf("unsupported amt boot target: '%s'", target)

	}

	// If we reach here, then the specified action is not supported.
	return fmt.Errorf("action '%s' is not supported for AMT boot target devices", action)
}
