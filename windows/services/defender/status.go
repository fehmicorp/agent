package defender

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
