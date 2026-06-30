package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/StackExchange/wmi"
	"github.com/fehmicorp/agent/windows/debug/logger"
	log "github.com/fehmicorp/agent/windows/debug/logger"
	"github.com/fehmicorp/agent/windows/services/metric"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	prevNetState      map[string]net.IOCountersStat
	lastNetworkSample time.Time
)

func (a *App) StartMetric(intervalMs int) {
	ticker := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-a.ctx.Done():
			return
		case <-ticker.C:
			cpuStr := metric.GetCPUUsage()
			ramStr := metric.GetRAMUsage()
			diskStr := metric.GetDiskUsage("C:")
			_, rxStr, err := metric.GetNetwork()
			netStr := "↓ 0B"
			if err == nil {
				netStr = fmt.Sprintf("↓ %s", rxStr)
			} else {
				log.Logger.Log("ERROR", "Network monitoring metrics update dropped: "+err.Error())
			}
			runtime.EventsEmit(a.ctx, "metrics_update", UsageStats{
				CPU:     cpuStr,
				RAM:     ramStr,
				Disk:    diskStr,
				Network: netStr,
			})
		}
	}
}

func getWmiDomain() string {
	var systemInfo []Win32_ComputerSystem
	query := "SELECT Domain, PartOfDomain FROM Win32_ComputerSystem"
	if err := wmi.Query(query, &systemInfo); err != nil || len(systemInfo) == 0 {
		return "Workgroup"
	}
	sys := systemInfo[0]
	if sys.Domain != "" {
		return sys.Domain
	}
	return "Workgroup"
}

func getExactOSName() string {
	info, err := host.Info()
	if err != nil {
		return "Unknown OS"
	}
	if info.Platform != "" {
		return info.Platform
	}
	return info.OS
}

func formatToTitleCase(s string) string {
	if s == "" {
		return ""
	}
	lower := strings.ToLower(s)
	r := []rune(lower)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func (a *App) DashboardUpdate() DeviceInfo {
	data, err := metric.GetSystemDeviceInfo()
	if err != nil {
		logger.Logger.Log("ERROR", "System device metrics extraction completed with partial degradation: "+err.Error())
		errStr := err.Error()
		if start := strings.Index(errStr, "["); start != -1 {
			if end := strings.LastIndex(errStr, "]"); end > start {
				rawErrors := errStr[start+1 : end]
				individualErrors := strings.Split(rawErrors, "; ")
				for _, singleErr := range individualErrors {
					if singleErr != "" {
						logger.Logger.Log("WARN", "Subsystem Failure Details -> "+singleErr)
					}
				}
			}
		}
	} else {
		return data
	}
}

// func GetSystemDeviceInfo() DeviceInfo {
// 	hostname, _ := os.Hostname()
// 	currentUser, _ := user.Current()
// 	osDisplay := getExactOSName()
// 	username := "Administrator"
// 	domain := formatToTitleCase(getWmiDomain())
// 	defender := getDefender()
// 	firewall := getFirewall()
// 	if currentUser != nil {
// 		rawUsername := currentUser.Username
// 		if strings.Contains(rawUsername, "\\") {
// 			parts := strings.Split(rawUsername, "\\")
// 			username = parts[len(parts)-1]
// 		} else {
// 			username = rawUsername
// 		}
// 	}

// 	Logger.Log("INFO", "Device info layout requested and built successfully")
// 	Logger.Log("SECURITY", fmt.Sprintf("Defender Overall: %s | Realtime: %s | Tamper: %s", defender.Status, defender.RealTimeProtection, defender.TamperProtection))
// 	Logger.Log("SECURITY", fmt.Sprintf("Firewall Overall: %s | Private: %s | Public: %s", firewall.Status, firewall.PrivateProfile, firewall.PublicProfile))
// 	return DeviceInfo{
// 		Hostname:        hostname,
// 		Domain:          domain,
// 		User:            username,
// 		OS:              osDisplay,
// 		AgentVersion:    "v1.0.1",
// 		WindowsDefender: defender.Status,
// 		Firewall:        firewall.Status,
// 		TPM:             "Enabled",
// 		BitLocker:       "Disabled",
// 	}
// }
