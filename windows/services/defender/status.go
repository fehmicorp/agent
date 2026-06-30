package defender

func GetDefenderStatus() (*DefenderStatus, error) {
	result := &DefenderStatus{
		Status:               "Disabled",
		SignatureVersion:     "0.0.0.0",
		SignatureLastUpdated: "Unknown",
		DefenderServiceState: "Stopped",
	}

	// Bubble up any underlying query issues if they happen
	if err := getPreferences(result); err != nil {
		return nil, err
	}
	if err := getSignatureDetails(result); err != nil {
		return nil, err
	}
	if err := checkServiceFallback(result); err != nil {
		return nil, err
	}

	evaluateAggregateHealth(result)

	return result, nil
}

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
