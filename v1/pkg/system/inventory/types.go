package inventory

import "time"

type InventoryInfo struct {
	CollectedAt time.Time `json:"collectedAt"`

	System   SystemInfo   `json:"system"`
	OS       OSInfo       `json:"os"`
	Hardware HardwareInfo `json:"hardware"`
	CPU      CPUInfo      `json:"cpu"`
	Memory   MemoryInfo   `json:"memory"`
	Storage  []DiskInfo   `json:"storage"`
	Network  []NICInfo    `json:"network"`

	Security SecurityInfo `json:"security"`

	Agent AgentInfo `json:"agent"`
}

type SystemInfo struct {
	// Unique identifiers
	Fingerprint string `json:"fingerprint"` // SHA-256 generated from multiple identifiers
	DeviceID    string `json:"deviceId"`    // Internal Fehmi persistent ID

	// Identity
	Hostname    string `json:"hostname"`
	Domain      string `json:"domain,omitempty"`
	Workgroup   string `json:"workgroup,omitempty"`
	User        string `json:"user"`
	DisplayName string `json:"displayName,omitempty"`

	// Hardware
	Manufacturer string `json:"manufacturer,omitempty"`
	Model        string `json:"model,omitempty"`

	// Vendor identifiers
	SerialNumber string `json:"serialNumber,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	BIOSSerial   string `json:"biosSerial,omitempty"`
	BoardSerial  string `json:"boardSerial,omitempty"`

	// Network
	PrimaryMAC string `json:"primaryMac,omitempty"`
}

type OSInfo struct {
	Platform     string `json:"platform"` // windows, linux, darwin
	Name         string `json:"name"`
	Version      string `json:"version"`
	Build        string `json:"build,omitempty"`
	Architecture string `json:"architecture"`
	Kernel       string `json:"kernel,omitempty"`

	InstallDate string `json:"installDate,omitempty"`
	BootTime    string `json:"bootTime,omitempty"`
	Uptime      uint64 `json:"uptime,omitempty"`
}

type HardwareInfo struct {
	BIOSVersion string `json:"biosVersion,omitempty"`
	BIOSVendor  string `json:"biosVendor,omitempty"`

	Motherboard string `json:"motherboard,omitempty"`

	Chassis string `json:"chassis,omitempty"`

	Battery bool `json:"battery"`
}

type CPUInfo struct {
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer,omitempty"`

	Cores           int `json:"cores"`
	LogicalCores    int `json:"logicalCores"`
	MaxClockMHz     int `json:"maxClockMHz,omitempty"`
	CurrentClockMHz int `json:"currentClockMHz,omitempty"`
}

type MemoryInfo struct {
	Installed uint64 `json:"installed"`
	Available uint64 `json:"available"`
	Used      uint64 `json:"used"`
	Slots     int    `json:"slots,omitempty"`
}

type DiskInfo struct {
	Name       string `json:"name"`
	FileSystem string `json:"fileSystem"`

	Total uint64 `json:"total"`
	Free  uint64 `json:"free"`
	Used  uint64 `json:"used"`

	SSD bool `json:"ssd,omitempty"`
}

type NICInfo struct {
	Name string `json:"name"`

	MAC string `json:"mac"`

	IPv4 string `json:"ipv4,omitempty"`
	IPv6 string `json:"ipv6,omitempty"`

	Gateway string `json:"gateway,omitempty"`

	DNS []string `json:"dns,omitempty"`

	Speed int64 `json:"speed,omitempty"`

	Wireless bool `json:"wireless"`
}

type SecurityInfo struct {
	Firewall bool `json:"firewall"`

	TPM bool `json:"tpm"`

	Encryption bool `json:"encryption"`

	WindowsDefender bool `json:"windowsDefender,omitempty"`

	BitLocker bool `json:"bitLocker,omitempty"`

	FileVault bool `json:"fileVault,omitempty"`

	SecureBoot bool `json:"secureBoot,omitempty"`

	Antivirus string `json:"antivirus,omitempty"`

	AntivirusEnabled bool `json:"antivirusEnabled"`
}

type AgentInfo struct {
	Version string `json:"version"`

	Build string `json:"build"`

	Online bool `json:"online"`

	LastSeen time.Time `json:"lastSeen"`
}
