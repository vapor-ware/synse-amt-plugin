package devices

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/scripts"
)

const (
	pxeTarget = "pxe"
	hdTarget  = "hd"
	cdTarget  = "cd"
)

// AmtBootTarget is the handler for setting an amt device's boot target
var AmtBootTarget = sdk.DeviceHandler{
	Name:  "boot_target",
	Write: bootTargetWrite,
}

// bootTargetWrite sets the amt boot target
func bootTargetWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Data

	// When writing to a boot_target device, we always expect there to be
	// raw data specified. If there isn't, we return an error.
	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'raw', but required")
	}

	if action == "target" {
		target := string(raw)

		switch strings.ToLower(target) {
		case pxeTarget, hdTarget, cdTarget:
			log.Infof("setting boot target to: %s", target)
		default:
			return fmt.Errorf("unsupported amt boot target: '%s'", target)
		}

		ip := fmt.Sprint(device.Data["ip"])
		pass := fmt.Sprint(device.Data["password"])

		cmd := scripts.NewCommand("python", "scripts/boot_target.py", ip, pass, target)
		err := cmd.Run()
		if err != nil {
			log.Errorf("error: %s", cmd.Stderr())
			return err
		}
		return nil
	}

	// If we reach here, then the specified action is not supported.
	return fmt.Errorf("action '%s' is not supported for AMT boot target devices", action)
}
