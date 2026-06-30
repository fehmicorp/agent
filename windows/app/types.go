package main

type DefenderDetails struct {
	Status             string // Overall Status: "Secure", "Action Required", or "Disabled"
	RealTimeProtection string // "Enabled" or "Disabled"
	TamperProtection   string // "Enabled" or "Disabled"
	BehaviorMonitoring string // "Enabled" or "Disabled"
	IOAVProtection     string // "Enabled" or "Disabled" (IE/Edge/Download scanning)
}

type FirewallDetails struct {
	Status         string // Overall Status: "Secure" or "Action Required"
	DomainProfile  string // "Enabled" or "Disabled"
	PrivateProfile string // "Enabled" or "Disabled"
	PublicProfile  string // "Enabled" or "Disabled"
}

// Native Adva	nced Firewall WMI mapping structure
type MSFT_NetFirewallProfileDetailed struct {
	Enabled bool   `wmi:"Enabled"`
	Profile uint16 `wmi:"Profile"` // 1: Domain, 2: Private, 4: Public
}

type Win32_PerfFormattedData_Amd64_HNSNetwork struct {
	// Alternatively, we use the standard Win32 firewall profile checks
}

type Win32_OperatingSystem struct {
	Caption                string
	TotalVisibleMemorySize uint64
	FreePhysicalMemory     uint64
}

type Win32_ComputerSystem struct {
	Domain       string
	PartOfDomain bool
}

type Win32_PerfFormattedData_Tcpip_NetworkInterface struct {
	Name                string
	BytesReceivedPerSec uint64
	BytesSentPerSec     uint64
}

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
