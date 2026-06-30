package defender

func evaluateAggregateHealth(status *DefenderStatus) {
	if !status.AntivirusEnabled || status.DefenderServiceState == "Stopped" {
		status.Status = "Disabled"
		return
	}

	if status.RealTimeProtection && status.TamperProtection {
		status.Status = "Secure"
	} else {
		status.Status = "Action Required"
	}
}
