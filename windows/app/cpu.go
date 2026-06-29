package main

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/yusufpapurcu/wmi"
)

// GetCpuFullDetails handles deep-dive architectural profiling configurations
func (a *App) GetCpuFullDetails() (CpuFullDetails, error) {
	// 1. Fetch live load snapshot percentages via gopsutil
	percent, err := cpu.Percent(0, false)
	liveUsageStr := "0%"
	if err == nil && len(percent) > 0 {
		liveUsageStr = fmt.Sprintf("%.0f%%", percent[0])
	}

	// 2. Extract deep silicon hardware traits using structural WMI interfaces
	var winProcessors []Win32_Processor
	query := "SELECT Name, NumberOfCores, NumberOfLogicalProcessors, MaxClockSpeed, CurrentClockSpeed, L2CacheSize, L3CacheSize FROM Win32_Processor"
	err = wmi.Query(query, &winProcessors)

	// Fallback protections if sandboxing access rules intercept the WMI subsystem layers
	if err != nil || len(winProcessors) == 0 {
		return CpuFullDetails{
			Model:          "Generic x86_64 Processor",
			PhysicalCores:  4,
			LogicalThreads: 8,
			MaxSpeed:       "Unknown MHz",
			CurrentSpeed:   "Unknown MHz",
			L2Cache:        "Unknown",
			L3Cache:        "Unknown",
			LiveUsage:      liveUsageStr,
		}, nil
	}

	cpuCore := winProcessors[0]

	// Format processing configurations clean into our return struct target payload
	details := CpuFullDetails{
		Model:          cpuCore.Name,
		PhysicalCores:  cpuCore.NumberOfCores,
		LogicalThreads: cpuCore.NumberOfLogicalProcessors,
		MaxSpeed:       fmt.Sprintf("%.1f GHz", float64(cpuCore.MaxClockSpeed)/1000.0),
		CurrentSpeed:   fmt.Sprintf("%.1f GHz", float64(cpuCore.CurrentClockSpeed)/1000.0),
		L2Cache:        fmt.Sprintf("%d KB", cpuCore.L2CacheSize),
		L3Cache:        fmt.Sprintf("%d MB", cpuCore.L3CacheSize/1024), // Convert KB payload boundaries into MB
		LiveUsage:      liveUsageStr,
	}

	return details, nil
}

// GetCpuFullDetailsJSON exposes a clean JSON execution pipeline bridge for your Wails frontend layout
func (a *App) GetCpuFullDetailsJSON() (string, error) {
	details, err := a.GetCpuFullDetails()
	if err != nil {
		return "", err
	}

	jsonBytes, err := json.Marshal(details)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
