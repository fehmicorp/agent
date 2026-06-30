package main

import (
	"fmt"

	"github.com/StackExchange/wmi"
)

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

type MSFT_MpComputerStatus struct {
	AntispywareEnabled            bool   `wmi:"AntispywareEnabled"`
	AntivirusEnabled              bool   `wmi:"AntivirusEnabled"`
	AntispywareSignatureVersion   string `wmi:"AntispywareSignatureVersion"`
	AntivirusSignatureLastUpdated string `wmi:"AntivirusSignatureLastUpdated"` // Fixed type to string
	QuickScanAge                  uint32 `wmi:"QuickScanAge"`
	FullScanAge                   uint32 `wmi:"FullScanAge"`
}

type Win32Service struct {
	State string `wmi:"State"`
}

func GetDefenderStatus() (*DefenderStatus, error) {
	result := &DefenderStatus{
		Status:               "Disabled",
		SignatureVersion:     "0.0.0.0",
		SignatureLastUpdated: "Unknown",
		DefenderServiceState: "Stopped",
	}

	// Collect Defender preferences
	if err := getPreferences(result); err != nil {
		return nil, err
	}

	// Collect signature information
	if err := getSignatureDetails(result); err != nil {
		return nil, err
	}

	// Verify Defender service state
	if err := checkServiceFallback(result); err != nil {
		return nil, err
	}

	// Compute overall health
	evaluateAggregateHealth(result)

	return result, nil
}

func evaluateAggregateHealth(status *DefenderStatus) {
	if status == nil {
		return
	}

	// Defender service is not running
	if status.DefenderServiceState != "Running" {
		status.Status = "Disabled"
		return
	}

	// Antivirus engine is disabled
	if !status.AntivirusEnabled {
		status.Status = "Disabled"
		return
	}

	// All major protections are enabled
	if status.RealTimeProtection &&
		status.BehaviorMonitoring &&
		status.IOAVProtection &&
		status.TamperProtection {
		status.Status = "Secure"
		return
	}

	// Defender is running but one or more protections are disabled
	status.Status = "Action Required"
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

func getPreferences(status *DefenderStatus) error {
	var prefs []MSFT_MpPreference
	query := "SELECT DisableRealtimeMonitoring, DisableBehaviorMonitoring, DisableIOAVProtection, EnableTamperProtection FROM MSFT_MpPreference"

	err := wmi.QueryNamespace(query, &prefs, "ROOT\\Microsoft\\Windows\\Defender")
	if err != nil {
		return err
	}

	if len(prefs) > 0 {
		p := prefs[0]
		status.RealTimeProtection = !p.DisableRealtimeMonitoring
		status.BehaviorMonitoring = !p.DisableBehaviorMonitoring
		status.IOAVProtection = !p.DisableIOAVProtection
		status.TamperProtection = (p.EnableTamperProtection == 4)
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
