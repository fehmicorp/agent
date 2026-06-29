package main

import "github.com/StackExchange/wmi"

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
