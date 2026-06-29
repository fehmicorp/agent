package metric

import (
	"fmt"
	"strings"

	"github.com/StackExchange/wmi"
	"github.com/fehmicorp/agent/windows/debug/logger"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type Win32_PerfFormattedData_Tcpip_NetworkInterface struct {
	Name                string
	BytesReceivedPerSec uint64
	BytesSentPerSec     uint64
}

func GetCPUUsage() string {
	cpuPercent, err := cpu.Percent(0, false)
	if err == nil && len(cpuPercent) > 0 {
		return fmt.Sprintf("%.0f%%", cpuPercent[0])
	}
	return "0%"
}

func GetRAMUsage() string {
	if vm, err := mem.VirtualMemory(); err == nil {
		return fmt.Sprintf("%.0f%%", vm.UsedPercent)
	} else {
		logger.Logger.Log("ERROR", "RAM metrics collection failed: "+err.Error())
	}
	return "0%"
}

func GetDiskUsage(drive string) string {
	d, err := disk.Usage(drive)
	if err != nil {
		logger.Logger.Log("ERROR", "Disk tracking error: "+err.Error())
		return "0%"
	}
	return fmt.Sprintf("%.0f%%", d.UsedPercent)
}

func GetWmiNetwork() (uint64, uint64, error) {
	var netAdapters []Win32_PerfFormattedData_Tcpip_NetworkInterface
	query := "SELECT Name, BytesReceivedPerSec, BytesSentPerSec FROM Win32_PerfFormattedData_Tcpip_NetworkInterface"
	if err := wmi.Query(query, &netAdapters); err != nil {
		return 0, 0, err
	}
	var rxTotal, txTotal uint64
	for _, adapter := range netAdapters {
		name := adapter.Name
		if contains(name, "Loopback") ||
			contains(name, "vEthernet") ||
			contains(name, "Virtual") ||
			contains(name, "Pseudo") ||
			contains(name, "Teredo") {
			continue
		}
		rxTotal += adapter.BytesReceivedPerSec
		txTotal += adapter.BytesSentPerSec
	}
	return rxTotal, txTotal, nil
}

func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func FormatSpeed(bytes float64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	var res string
	switch {
	case bytes >= GB:
		res = fmt.Sprintf("%.2fG", bytes/GB)
	case bytes >= MB:
		res = fmt.Sprintf("%.2fM", bytes/MB)
	case bytes >= KB:
		res = fmt.Sprintf("%.1fK", bytes/KB)
	default:
		return fmt.Sprintf("%.0fB", bytes)
	}
	res = strings.TrimSuffix(res, ".00G")
	res = strings.TrimSuffix(res, ".00M")
	res = strings.TrimSuffix(res, ".0K")
	if strings.Contains(res, ".") && (strings.HasSuffix(res, "G") || strings.HasSuffix(res, "M")) {
		res = strings.TrimSuffix(res, "0G") + "G"
		res = strings.TrimSuffix(res, "0M") + "M"
		res = strings.Replace(res, ".G", "G", 1)
		res = strings.Replace(res, ".M", "M", 1)
	}
	return res
}
