package inventory

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/fehmicorp/agent/v1/cmd/appinfo"
)

var deviceID string

func DeviceID(inv *InventoryInfo, product string) (string, error) {

	if deviceID != "" {
		return deviceID, nil
	}

	platform := "unk"

	switch runtime.GOOS {
	case "windows":
		platform = "win"
	case "linux":
		platform = "linux"
	case "darwin":
		platform = "mac"
	}

	version := "v1"

	if appinfo.Current.Version != "" {
		parts := strings.Split(appinfo.Current.Version, ".")
		if len(parts) > 0 {
			version = "v" + parts[0]
		}
	}

	mac := strings.ReplaceAll(inv.System.PrimaryMAC, ":", "")
	mac = strings.ReplaceAll(mac, "-", "")

	uuid := strings.ReplaceAll(inv.System.UUID, "-", "")

	if len(uuid) >= 12 {
		uuid = uuid[len(uuid)-12:]
	}

	deviceID = fmt.Sprintf(
		"%s%s-%s-%s-%s",
		product,
		version,
		platform,
		mac,
		strings.ToLower(uuid),
	)

	return deviceID, nil
}
