package shared

import (
	"log"
	"os"
)

// DeviceContain ...
func DeviceContain(target string) string {
	devices := []string{
		DataObjectLight,
		DataObjectAircon,
		DataObjectProjector,
		DataObjectMonitor,
		DataObjectTV,
	}
	for _, device := range devices {
		if device == target {
			return target
		}
	}

	return ""
}

// CheckDirectory ...
func CheckDirectory(path string) {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		if err := os.Mkdir(path, 0750); err != nil {
			log.Fatal(err)
		}
	}
}
