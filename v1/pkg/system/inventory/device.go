package inventory

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/fehmicorp/agent/v1/cmd/appinfo"
	"github.com/google/uuid"
)

var deviceID string

func DeviceID() (string, error) {
	if deviceID != "" {
		return deviceID, nil
	}
	path := filepath.Join(
		appinfo.Dir.Config,
		"device.id",
	)
	if b, err := os.ReadFile(path); err == nil {

		deviceID = strings.TrimSpace(string(b))

		return deviceID, nil
	}
	id := uuid.NewString()
	if err := os.MkdirAll(appinfo.Dir.Config, 0755); err != nil {
		return "", err
	}
	if err := os.WriteFile(path, []byte(id), 0644); err != nil {
		return "", err
	}
	deviceID = id

	return deviceID, nil
}
