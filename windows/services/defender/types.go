package main

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

type MSFT_MpPreference struct {
	DisableRealtimeMonitoring bool  `wmi:"DisableRealtimeMonitoring"`
	DisableBehaviorMonitoring bool  `wmi:"DisableBehaviorMonitoring"`
	DisableIOAVProtection     bool  `wmi:"DisableIOAVProtection"`
	EnableTamperProtection    uint8 `wmi:"EnableTamperProtection"` // 0: Disabled, 4: Enabled
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
