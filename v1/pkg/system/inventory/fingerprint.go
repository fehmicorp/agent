package inventory

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func GenerateFingerprint(inv *InventoryInfo) string {
	s := strings.Join([]string{
		inv.System.UUID,
		inv.System.SerialNumber,
		inv.System.BoardSerial,
		inv.System.BIOSSerial,
		inv.System.PrimaryMAC,
		inv.System.Manufacturer,
		inv.System.Model,
		inv.CPU.Model,
		inv.OS.Platform,
		inv.OS.Architecture,
	}, "|")
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}
