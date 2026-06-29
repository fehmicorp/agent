package main

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/yusufpapurcu/wmi"
)

// GetMemoryFullDetails collects metrics and layout, returning a raw Go Struct for internal performance
func (a *App) GetMemoryFullDetails() (MemoryDetails, error) {
	vMem, err := mem.VirtualMemory()
	if err != nil {
		return MemoryDetails{}, err
	}

	var winMems []Win32_PhysicalMemory
	query := "SELECT DeviceLocator, Capacity FROM Win32_PhysicalMemory"
	_ = wmi.Query(query, &winMems)

	slotsUsed := len(winMems)
	totalSlots := slotsUsed
	if totalSlots == 0 {
		totalSlots = 2
		slotsUsed = 1
	}

	totalGB := float64(vMem.Total) / (1024 * 1024 * 1024)
	freeGB := float64(vMem.Available) / (1024 * 1024 * 1024)

	return MemoryDetails{
		Used:       fmt.Sprintf("%.0f%%", vMem.UsedPercent),
		Capacity:   fmt.Sprintf("%.0fGB", totalGB),
		TotalSlots: totalSlots,
		SlotUsed:   slotsUsed,
		Free:       fmt.Sprintf("%.1fGB", freeGB),
	}, nil
}

// GetMemoryFullDetailsJSON acts as the Wails frontend runtime handler bridge
func (a *App) GetMemoryFullDetailsJSON() (string, error) {
	details, err := a.GetMemoryFullDetails()
	if err != nil {
		return "", err
	}
	jsonBytes, err := json.Marshal(details)
	return string(jsonBytes), err
}

// GetRamDetails fetches comprehensive characteristics for a specific slot index
func (a *App) GetRamDetails(slotIndex int) (string, error) {
	var winMems []Win32_PhysicalMemory
	query := "SELECT DeviceLocator, Speed, FormFactor, Capacity FROM Win32_PhysicalMemory"
	err := wmi.Query(query, &winMems)
	if err != nil {
		return "", err
	}

	if slotIndex < 0 || slotIndex >= len(winMems) {
		return "{}", fmt.Errorf("requested slot index %d is empty or unavailable", slotIndex)
	}

	targetStick := winMems[slotIndex]

	formFactorStr := "Unknown"
	switch targetStick.FormFactor {
	case 8:
		formFactorStr = "SO-DIMM (Laptop)"
	case 12:
		formFactorStr = "DIMM (Desktop)"
	}

	stickGB := float64(targetStick.Capacity) / (1024 * 1024 * 1024)

	spec := RamSpecification{
		Speed:      fmt.Sprintf("%d MT/s", targetStick.Speed),
		FormFactor: formFactorStr,
		Capacity:   fmt.Sprintf("%.0f GB", math.Round(stickGB)),
		Locator:    targetStick.DeviceLocator,
	}

	jsonBytes, err := json.Marshal(spec)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
