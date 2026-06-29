package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"
	"unicode"

	"github.com/StackExchange/wmi"
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
			netStr := metric.GetNetwork()
			runtime.EventsEmit(a.ctx, "metrics_update", UsageStats{
				CPU:     cpuStr,
				RAM:     ramStr,
				Disk:    diskStr,
				Network: netStr,
			})
		}
	}
}

// func getCPUUsage() string {
// 	cpuPercent, err := cpu.Percent(0, false)
// 	if err == nil && len(cpuPercent) > 0 {
// 		return fmt.Sprintf("%.0f%%", cpuPercent[0])
// 	}

// 	if err != nil {
// 		Logger.Log("ERROR", "CPU metrics collection failed: "+err.Error())
// 	}
// 	return "0%"
// }

// func getRAMUsage() string {
// 	if vm, err := mem.VirtualMemory(); err == nil {
// 		return fmt.Sprintf("%.0f%%", vm.UsedPercent)
// 	} else {
// 		Logger.Log("ERROR", "RAM metrics collection failed: "+err.Error())
// 	}
// 	return "0%"
// }

// func getDiskUsage(drive string) string {
// 	d, err := disk.Usage(drive)
// 	if err != nil {
// 		Logger.Log("ERROR", "Disk tracking error: "+err.Error())
// 		return "0%"
// 	}
// 	return fmt.Sprintf("%.0f%%", d.UsedPercent)
// }

// func getWmiNetwork() (uint64, uint64, error) {
// 	var netAdapters []Win32_PerfFormattedData_Tcpip_NetworkInterface
// 	query := "SELECT Name, BytesReceivedPerSec, BytesSentPerSec FROM Win32_PerfFormattedData_Tcpip_NetworkInterface"
// 	if err := wmi.Query(query, &netAdapters); err != nil {
// 		return 0, 0, err
// 	}
// 	var rxTotal, txTotal uint64
// 	for _, adapter := range netAdapters {
// 		name := adapter.Name
// 		if contains(name, "Loopback") ||
// 			contains(name, "vEthernet") ||
// 			contains(name, "Virtual") ||
// 			contains(name, "Pseudo") ||
// 			contains(name, "Teredo") {
// 			continue
// 		}
// 		rxTotal += adapter.BytesReceivedPerSec
// 		txTotal += adapter.BytesSentPerSec
// 	}
// 	return rxTotal, txTotal, nil
// }

// func contains(s, substr string) bool {
// 	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
// }

// func formatSpeed(bytes float64) string {
// 	const (
// 		KB = 1024
// 		MB = KB * 1024
// 		GB = MB * 1024
// 	)

// 	var res string
// 	switch {
// 	case bytes >= GB:
// 		res = fmt.Sprintf("%.2fG", bytes/GB)
// 	case bytes >= MB:
// 		res = fmt.Sprintf("%.2fM", bytes/MB)
// 	case bytes >= KB:
// 		res = fmt.Sprintf("%.1fK", bytes/KB)
// 	default:
// 		return fmt.Sprintf("%.0fB", bytes)
// 	}
// 	res = strings.TrimSuffix(res, ".00G")
// 	res = strings.TrimSuffix(res, ".00M")
// 	res = strings.TrimSuffix(res, ".0K")
// 	if strings.Contains(res, ".") && (strings.HasSuffix(res, "G") || strings.HasSuffix(res, "M")) {
// 		res = strings.TrimSuffix(res, "0G") + "G"
// 		res = strings.TrimSuffix(res, "0M") + "M"
// 		res = strings.Replace(res, ".G", "G", 1)
// 		res = strings.Replace(res, ".M", "M", 1)
// 	}
// 	return res
// }

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

func (a *App) GetSystemDeviceInfo() DeviceInfo {
	hostname, _ := os.Hostname()
	currentUser, _ := user.Current()
	osDisplay := getExactOSName()
	username := "Administrator"
	domain := formatToTitleCase(getWmiDomain())
	defender := getDefender()
	firewall := getFirewall()
	if currentUser != nil {
		rawUsername := currentUser.Username
		if strings.Contains(rawUsername, "\\") {
			parts := strings.Split(rawUsername, "\\")
			username = parts[len(parts)-1]
		} else {
			username = rawUsername
		}
	}

	Logger.Log("INFO", "Device info layout requested and built successfully")
	Logger.Log("SECURITY", fmt.Sprintf("Defender Overall: %s | Realtime: %s | Tamper: %s", defender.Status, defender.RealTimeProtection, defender.TamperProtection))
	Logger.Log("SECURITY", fmt.Sprintf("Firewall Overall: %s | Private: %s | Public: %s", firewall.Status, firewall.PrivateProfile, firewall.PublicProfile))
	return DeviceInfo{
		Hostname:        hostname,
		Domain:          domain,
		User:            username,
		OS:              osDisplay,
		AgentVersion:    "v1.0.1",
		WindowsDefender: defender.Status,
		Firewall:        firewall.Status,
		TPM:             "Enabled",
		BitLocker:       "Disabled",
	}
}
