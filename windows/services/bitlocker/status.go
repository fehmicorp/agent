package bitlocker

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

type BitLockerStatus struct {
	Drive             string `json:"drive"`
	ProtectionEnabled bool   `json:"protection_enabled"`
	Encrypted         bool   `json:"encrypted"`
	Locked            bool   `json:"locked"`
	EncryptionMethod  string `json:"encryption_method"`
	ProtectionStatus  string `json:"protection_status"`
	VolumeStatus      string `json:"volume_status"`
	LockStatus        string `json:"lock_status"`
}

type psBitLocker struct {
	MountPoint       string `json:"MountPoint"`
	ProtectionStatus int    `json:"ProtectionStatus"`
	VolumeStatus     int    `json:"VolumeStatus"`
	LockStatus       int    `json:"LockStatus"`
	EncryptionMethod int    `json:"EncryptionMethod"`
}

func GetBitLockerStatus() ([]BitLockerStatus, error) {

	script := "Get-BitLockerVolume | Select-Object MountPoint,ProtectionStatus,VolumeStatus,LockStatus,EncryptionMethod | ConvertTo-Json -Compress"

	cmd := exec.Command(
		"powershell.exe",
		"-NoLogo",
		"-NoProfile",
		"-NonInteractive",
		"-ExecutionPolicy", "Bypass",
		"-Command",
		script,
	)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: windows.CREATE_NO_WINDOW,
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, string(out))
	}

	// PowerShell returns either an object or an array.
	var raw []psBitLocker

	if err := json.Unmarshal(out, &raw); err != nil {

		var single psBitLocker

		if err2 := json.Unmarshal(out, &single); err2 != nil {
			return nil, err
		}

		raw = []psBitLocker{single}
	}

	result := make([]BitLockerStatus, 0, len(raw))

	for _, v := range raw {

		result = append(result, BitLockerStatus{
			Drive:             v.MountPoint,
			ProtectionEnabled: v.ProtectionStatus == 1,
			Encrypted:         v.VolumeStatus == 1,
			Locked:            v.LockStatus == 1,
			EncryptionMethod:  encryptionMethod(v.EncryptionMethod),
			ProtectionStatus:  protectionStatus(v.ProtectionStatus),
			VolumeStatus:      volumeStatus(v.VolumeStatus),
			LockStatus:        lockStatus(v.LockStatus),
		})
	}

	return result, nil
}

func protectionStatus(v int) string {

	switch v {

	case 0:
		return "Protection Off"

	case 1:
		return "Protection On"

	default:
		return "Unknown"
	}
}

func volumeStatus(v int) string {

	switch v {

	case 0:
		return "Fully Decrypted"

	case 1:
		return "Fully Encrypted"

	case 2:
		return "Encryption In Progress"

	case 3:
		return "Decryption In Progress"

	case 4:
		return "Encryption Paused"

	case 5:
		return "Decryption Paused"

	default:
		return "Unknown"
	}
}

func lockStatus(v int) string {

	switch v {

	case 0:
		return "Unlocked"

	case 1:
		return "Locked"

	default:
		return "Unknown"
	}
}

func encryptionMethod(v int) string {

	switch v {

	case 0:
		return "None"

	case 1:
		return "AES-128"

	case 2:
		return "AES-256"

	case 3:
		return "AES-128 Diffuser"

	case 4:
		return "AES-256 Diffuser"

	case 5:
		return "Hardware Encryption"

	case 6:
		return "XTS-AES 128"

	case 7:
		return "XTS-AES 256"

	default:
		return "Unknown"
	}
}
