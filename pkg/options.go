package pkg

import "fmt"

// deviceIdentifier defines the AMT-specific way of uniquely identifying a device
// through its device configuration.
func deviceIdentifier(data map[string]interface{}) string {
	return fmt.Sprint(data["ip"])
}
