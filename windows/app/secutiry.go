package main

import "github.com/StackExchange/wmi"

func getFirewall() FirewallDetails {
	details := FirewallDetails{
		Status:         "Action Required",
		DomainProfile:  "Disabled",
		PrivateProfile: "Disabled",
		PublicProfile:  "Disabled",
	}

	var profiles []MSFT_NetFirewallProfileDetailed
	query := "SELECT Enabled, Profile FROM MSFT_NetFirewallProfile"

	err := wmi.QueryNamespace(query, &profiles, "ROOT\\Microsoft\\Windows\\Firewall")
	if err == nil && len(profiles) > 0 {
		for _, prof := range profiles {
			stateStr := "Disabled"
			if prof.Enabled {
				stateStr = "Enabled"
			}

			// Map individual engine profile instances
			switch prof.Profile {
			case 1:
				details.DomainProfile = stateStr
			case 2:
				details.PrivateProfile = stateStr
			case 4:
				details.PublicProfile = stateStr
			}
		}

		// Secure if your primary operational surface blocks are active
		if details.PrivateProfile == "Enabled" || details.PublicProfile == "Enabled" {
			details.Status = "Secure"
		}
	} else {
		// Fallback service engine check
		var services []map[string]interface{}
		if err := wmi.Query("SELECT State FROM Win32_Service WHERE Name='MpsSvc'", &services); err == nil && len(services) > 0 {
			if services[0]["State"].(string) == "Running" {
				details.Status = "Secure"
				details.PrivateProfile = "Enabled"
				details.PublicProfile = "Enabled"
			}
		}
	}

	return details
}

func getDefender() DefenderDetails {
	details := DefenderDetails{
		Status:             "Disabled",
		RealTimeProtection: "Disabled",
		TamperProtection:   "Disabled",
		BehaviorMonitoring: "Disabled",
		IOAVProtection:     "Disabled",
	}

	var prefs []MSFT_MpPreference
	query := "SELECT DisableRealtimeMonitoring, DisableBehaviorMonitoring, DisableIOAVProtection, EnableTamperProtection FROM MSFT_MpPreference"

	// Querying the explicit Windows Defender Engine configuration namespace
	err := wmi.QueryNamespace(query, &prefs, "ROOT\\Microsoft\\Windows\\Defender")
	if err == nil && len(prefs) > 0 {
		p := prefs[0]

		if !p.DisableRealtimeMonitoring {
			details.RealTimeProtection = "Enabled"
		}
		if !p.DisableBehaviorMonitoring {
			details.BehaviorMonitoring = "Enabled"
		}
		if !p.DisableIOAVProtection {
			details.IOAVProtection = "Enabled"
		}
		if p.EnableTamperProtection == 4 {
			details.TamperProtection = "Enabled"
		}

		// Calculate an aggregation health score status
		if details.RealTimeProtection == "Enabled" && details.TamperProtection == "Enabled" {
			details.Status = "Secure"
		} else if details.RealTimeProtection == "Enabled" {
			details.Status = "Action Required" // Realtime is up but unprotected configuration settings exist
		}
	} else {
		// Non-admin or broken WMI fallback validation block
		var services []map[string]interface{}
		if err := wmi.Query("SELECT State FROM Win32_Service WHERE Name='WinDefend'", &services); err == nil && len(services) > 0 {
			if services[0]["State"].(string) == "Running" {
				details.Status = "Secure"
				details.RealTimeProtection = "Enabled"
			}
		}
	}

	return details
}
