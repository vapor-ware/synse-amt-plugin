package outputs

import "github.com/vapor-ware/synse-sdk/sdk"

var (
	// PowerState is the output type for power state readings (on/off).
	PowerState = sdk.OutputType{
		Name: "power.state",
	}

	// BootTarget is the output type for boot target settings.
	BootTarget = sdk.OutputType{
		Name: "boot.target",
	}
)
