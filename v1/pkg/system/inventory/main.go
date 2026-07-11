package inventory

import (
	"encoding/json"
	"time"

	"github.com/fehmicorp/agent/v1/res/logger"
)

func TestRun() {
	inv := &InventoryInfo{}
	inv.CollectedAt = time.Now().UTC()

	if err := collectSystem(&inv.System); err != nil {
		logger.Error("Failed to collect system info: %v", err)
		return
	}

	// Generate or load the persistent Device ID
	id, err := DeviceID(inv, "ag")
	if err != nil {
		logger.Error("Failed to get device ID: %v", err)
		return
	}
	inv.System.DeviceID = id

	// Generate fingerprint after system information is populated
	inv.System.Fingerprint = GenerateFingerprint(inv)

	data, err := json.MarshalIndent(inv.System, "", "  ")
	if err != nil {
		logger.Error("Failed to marshal system info: %v", err)
		return
	}

	logger.Info("%s", string(data))
}

func Scan() (*InventoryInfo, error) {
	inv := &InventoryInfo{}
	inv.CollectedAt = time.Now().UTC()
	if err := collectSystem(&inv.System); err != nil {
		return nil, err
	}
	if err := collectOS(&inv.OS); err != nil {
		return nil, err
	}
	if err := collectHardware(&inv.Hardware); err != nil {
		return nil, err
	}
	if err := collectCPU(&inv.CPU); err != nil {
		return nil, err
	}
	// if err := collectMemory(&inv.Memory); err != nil {
	// 	return nil, err
	// }
	// if err := collectDisk(&inv.Storage); err != nil {
	// 	return nil, err
	// }
	// if err := collectNetwork(&inv.Network); err != nil {
	// 	return nil, err
	// }
	// if err := collectSecurity(&inv.Security); err != nil {
	// 	return nil, err
	// }
	// inv.System.Fingerprint = GenerateFingerprint(inv)
	// id, err := DeviceID()
	// if err != nil {
	// 	return nil, err
	// }
	// inv.System.DeviceID = id
	return inv, nil
}
