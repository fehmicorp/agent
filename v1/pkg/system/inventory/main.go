package inventory

import "time"

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
	if err := collectMemory(&inv.Memory); err != nil {
		return nil, err
	}
	if err := collectDisk(&inv.Storage); err != nil {
		return nil, err
	}
	if err := collectNetwork(&inv.Network); err != nil {
		return nil, err
	}
	if err := collectSecurity(&inv.Security); err != nil {
		return nil, err
	}
	inv.System.Fingerprint = GenerateFingerprint(inv)
	id, err := DeviceID()
	if err != nil {
		return nil, err
	}
	inv.System.DeviceID = id
	return inv, nil
}
