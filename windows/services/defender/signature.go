package defender

import (
	"fmt"

	"github.com/StackExchange/wmi"
)

func getSignatureDetails(status *DefenderStatus) error {
	var compStatus []MSFT_MpComputerStatus
	query := "SELECT AntispywareEnabled, AntivirusEnabled, AntispywareSignatureVersion, AntivirusSignatureLastUpdated, QuickScanAge, FullScanAge FROM MSFT_MpComputerStatus"

	err := wmi.QueryNamespace(query, &compStatus, "ROOT\\Microsoft\\Windows\\Defender")
	if err != nil {
		return err
	}

	if len(compStatus) > 0 {
		c := compStatus[0]
		status.AntispywareEnabled = c.AntispywareEnabled
		status.AntivirusEnabled = c.AntivirusEnabled
		status.SignatureVersion = c.AntispywareSignatureVersion

		// Parse WMI CIM_DateTime string safely (e.g., "20260629...") into readable output
		if len(c.AntivirusSignatureLastUpdated) >= 14 {
			raw := c.AntivirusSignatureLastUpdated
			status.SignatureLastUpdated = fmt.Sprintf("%s-%s-%s %s:%s:%s",
				raw[0:4], raw[4:6], raw[6:8], raw[8:10], raw[10:12], raw[12:14])
		} else {
			status.SignatureLastUpdated = c.AntivirusSignatureLastUpdated
		}

		status.QuickScanAge = c.QuickScanAge
		status.FullScanAge = c.FullScanAge
	}
	return nil
}
