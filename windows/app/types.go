package main

type DeviceInfo struct {
	Hostname        string `json:"hostname"`
	Domain          string `json:"domain"`
	User            string `json:"user"`
	OS              string `json:"os"`
	AgentVersion    string `json:"agentVersion"`
	WindowsDefender string `json:"windowsDefender"`
	Firewall        string `json:"firewall"`
	TPM             string `json:"tpm"`
	BitLocker       string `json:"bitLocker"`
}

type UsageStats struct {
	CPU     string `json:"cpu"`
	RAM     string `json:"ram"`
	Disk    string `json:"disk"`
	Network string `json:"network"`
}

type NetworkSpeed struct {
	Download string `json:"download"`
	Upload   string `json:"upload"`
	RxBytes  uint64 `json:"rxBytes"`
	TxBytes  uint64 `json:"txBytes"`
}

// Win32_Processor maps high-level hardware layout schemas from Windows WMI core engine
type Win32_Processor struct {
	Name                      string // Manufacturer model description tag (e.g., "Intel Core i7-12700H")
	NumberOfCores             uint32 // Physical layout cores
	NumberOfLogicalProcessors uint32 // SMT Hyperthreading allocation slots
	MaxClockSpeed             uint32 // Factory target clock rating speed in MHz
	CurrentClockSpeed         uint32 // Live clock speed frequency in MHz
	L2CacheSize               uint32 // Level 2 caching capability in KB
	L3CacheSize               uint32 // Level 3 caching capability in KB
}

// CpuFullDetails defines the structural layout returned for your system diagnostics
type CpuFullDetails struct {
	Model          string `json:"model"`
	PhysicalCores  uint32 `json:"physicalCores"`
	LogicalThreads uint32 `json:"logicalThreads"`
	MaxSpeed       string `json:"maxSpeed"`
	CurrentSpeed   string `json:"currentSpeed"`
	L2Cache        string `json:"l2Cache"`
	L3Cache        string `json:"l3Cache"`
	LiveUsage      string `json:"liveUsage"`
}

type Win32_PhysicalMemory struct {
	DeviceLocator string // Slot identification tag (e.g., "DIMM 0")
	Speed         uint32 // Clock frequency speed in MT/s
	FormFactor    uint16 // Enclosure hardware identifier
	Capacity      uint64 // Physical volume per stick in bytes
}

type MemoryDetails struct {
	Used       string `json:"used"`
	Capacity   string `json:"capacity"`
	TotalSlots int    `json:"totalslots"`
	SlotUsed   int    `json:"slotused"`
	Free       string `json:"free"`
}

type RamSpecification struct {
	Speed      string `json:"speed"`
	FormFactor string `json:"formfactor"`
	Capacity   string `json:"capacity"`
	Locator    string `json:"locator"`
}
