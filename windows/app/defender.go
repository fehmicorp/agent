package main

import (
	"fmt"

	"github.com/StackExchange/wmi"
)

// DefenderStatus represents the aggregate state of Windows Defender telemetry
type DefenderStatus struct {
	Status               string `json:"status"` // "Secure", "Action Required", or "Disabled"
	RealTimeProtection   bool   `json:"real_time_protection"`
	BehaviorMonitoring   bool   `json:"behavior_monitoring"`
	IOAVProtection       bool   `json:"ioav_protection"`
	TamperProtection     bool   `json:"tamper_protection"`
	AntispywareEnabled   bool   `json:"antispyware_enabled"`
	AntivirusEnabled     bool   `json:"antivirus_enabled"`
	SignatureVersion     string `json:"signature_version"`
	SignatureLastUpdated string `json:"signature_last_updated"`
	QuickScanAge         uint32 `json:"quick_scan_age_days"`
	FullScanAge          uint32 `json:"full_scan_age_days"`
	DefenderServiceState string `json:"defender_service_state"`
}

// MSFT_MpPreference handles raw policy configurations from Windows Defender
type MSFT_MpPreference struct {
	DisableRealtimeMonitoring bool `wmi:"DisableRealtimeMonitoring"`
	DisableBehaviorMonitoring bool `wmi:"DisableBehaviorMonitoring"`
	DisableIOAVProtection     bool `wmi:"DisableIOAVProtection"`
}

// MSFT_MpComputerStatus maps native WMI engine protection health states
type MSFT_MpComputerStatus struct {
	AntispywareEnabled            bool   `wmi:"AntispywareEnabled"`
	AntivirusEnabled              bool   `wmi:"AntivirusEnabled"`
	AntispywareSignatureVersion   string `wmi:"AntispywareSignatureVersion"`
	AntivirusSignatureLastUpdated string `wmi:"AntivirusSignatureLastUpdated"`
	QuickScanAge                  uint32 `wmi:"QuickScanAge"`
	FullScanAge                   uint32 `wmi:"FullScanAge"`
}

// Win32Service reads the execution status lifecycle of underlying Windows NT services
type Win32Service struct {
	State string `wmi:"State"`
}

// GetDefenderStatus aggregates all subsystems to calculate runtime system endpoint safety posture
func GetDefenderStatus() (*DefenderStatus, error) {
	result := &DefenderStatus{
		Status:               "Disabled",
		SignatureVersion:     "0.0.0.0",
		SignatureLastUpdated: "Unknown",
		DefenderServiceState: "Stopped",
	}

	// 1. Collect policy preferences
	if err := getPreferences(result); err != nil {
		return nil, err
	}

	// 2. Collect engine signature telemetry updates
	if err := getSignatureDetails(result); err != nil {
		return nil, err
	}

	// 3. Verify underlying service lifecycle state
	if err := checkServiceFallback(result); err != nil {
		return nil, err
	}

	// 4. Evaluate consolidated absolute health context
	evaluateAggregateHealth(result)

	return result, nil
}

func getPreferences(status *DefenderStatus) error {
	var prefs []MSFT_MpPreference

	err := wmi.QueryNamespace(
		`SELECT
			DisableRealtimeMonitoring,
			DisableBehaviorMonitoring,
			DisableIOAVProtection
		FROM MSFT_MpPreference`,
		&prefs,
		`ROOT\Microsoft\Windows\Defender`,
	)
	if err != nil {
		return err
	}

	if len(prefs) == 0 {
		return nil
	}

	p := prefs[0]

	status.RealTimeProtection = !p.DisableRealtimeMonitoring
	status.BehaviorMonitoring = !p.DisableBehaviorMonitoring
	status.IOAVProtection = !p.DisableIOAVProtection
	tpm, err := GetTPMStatus()
	tpmStatus := false
	if err == nil && tpm != nil {
		tpmStatus =
			tpm.Present &&
				tpm.Ready &&
				tpm.Enabled &&
				tpm.Activated
	}
	status.TamperProtection = tpmStatus

	return nil
}

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

		// Parse WMI high-precision Timestamp format safely (e.g., YYYYMMDDHHMMSS...)
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

func checkServiceFallback(status *DefenderStatus) error {
	var services []Win32Service
	query := "SELECT State FROM Win32_Service WHERE Name='WinDefend'"

	err := wmi.Query(query, &services)
	if err != nil {
		status.DefenderServiceState = "Unknown"
		return err
	}

	if len(services) > 0 {
		status.DefenderServiceState = services[0].State
		if status.DefenderServiceState == "Running" && status.Status == "Disabled" {
			status.Status = "Secure"
			status.RealTimeProtection = true
		}
	} else {
		status.DefenderServiceState = "Unknown"
	}
	return nil
}

func evaluateAggregateHealth(status *DefenderStatus) {
	if status == nil {
		return
	}

	if status.DefenderServiceState != "Running" {
		status.Status = "Disabled"
		return
	}

	if !status.AntivirusEnabled {
		status.Status = "Disabled"
		return
	}

	if status.RealTimeProtection &&
		status.BehaviorMonitoring &&
		status.IOAVProtection &&
		status.TamperProtection {
		status.Status = "Secure"
		return
	}

	status.Status = "Action Required"
}
