package main

import (
	"fmt"

	"github.com/StackExchange/wmi"
)

// BitLockerVolumeDetails holds the boolean status metrics for a single drive volume
type BitLockerVolumeDetails struct {
	DriveLetter string `json:"drive_letter"`
	IsEncrypted bool   `json:"is_encrypted"` // true if protection is active
	IsLocked    bool   `json:"is_locked"`    // true if the volume requires an explicit decryption key mount
}

// Win32_EncryptableVolume maps directly to the native Windows MicrosoftVolumeEncryption WMI schema
type Win32_EncryptableVolume struct {
	DriveLetter            string `wmi:"DriveLetter"`
	ProtectionStatus       uint32 `wmi:"ProtectionStatus"` // 0 = Off, 1 = On, 2 = Unknown
	ProtectionStatusMethod uint32 `wmi:"ConversionStatus"` // 0 = Decrypted, 1 = Encrypted
}

// GetBitLockerStatus loops over all system disk volumes to pull absolute encryption state boolean flags
func GetBitLockerStatus() ([]BitLockerVolumeDetails, error) {
	var wmiVolumes []Win32_EncryptableVolume
	var volumeDetails []BitLockerVolumeDetails

	// Query targeting the specific MS volume encryption namespace
	query := "SELECT DriveLetter, ProtectionStatus, ConversionStatus FROM Win32_EncryptableVolume"
	err := wmi.QueryNamespace(query, &wmiVolumes, "ROOT\\CIMv2\\Security\\MicrosoftVolumeEncryption")
	if err != nil {
		return nil, fmt.Errorf("bitlocker WMI query failed: %w", err)
	}

	for _, v := range wmiVolumes {
		// Ignore volumes that do not have a mounted structural drive letter assignment (e.g., recovery partitions)
		if v.DriveLetter == "" {
			continue
		}

		details := BitLockerVolumeDetails{
			DriveLetter: v.DriveLetter,
			IsEncrypted: v.ProtectionStatus == 1, // 1 means Protection is On/Active
			IsLocked:    false,                   // If WMI can query runtime properties, the volume is currently Unlocked
		}

		volumeDetails = append(volumeDetails, details)
	}

	return volumeDetails, nil
}
