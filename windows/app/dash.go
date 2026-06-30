package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fehmicorp/agent/windows/debug/logger"
	"github.com/fehmicorp/agent/windows/services/firewall"
	"github.com/fehmicorp/agent/windows/services/metric"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	prevNetState      map[string]net.IOCountersStat
	lastNetworkSample time.Time
)

type UsageStats struct {
	CPU     string `json:"cpu"`
	RAM     string `json:"ram"`
	Disk    string `json:"disk"`
	Network string `json:"network"`
}

func (a *App) fetchAndEmit() {
	cpuStr := metric.GetCPUUsage()
	ramStr := metric.GetRAMUsage()
	diskStr := metric.GetDiskUsage("C:")
	_, rxStr, err := metric.GetNetwork()
	netStr := "↓ 0B"
	if err == nil {
		netStr = fmt.Sprintf("↓ %s", rxStr)
	} else {
		logger.Logger.Log("ERROR", "Network monitoring metrics update dropped: "+err.Error())
	}

	runtime.EventsEmit(a.ctx, "metrics_update", UsageStats{
		CPU:     cpuStr,
		RAM:     ramStr,
		Disk:    diskStr,
		Network: netStr,
	})
}

func (a *App) StartMetric(intervalMs int) {
	a.fetchAndEmit()
	ticker := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-a.ctx.Done():
			return
		case <-ticker.C:
			a.fetchAndEmit()
		}
	}
}

func GetSystemDeviceInfo() (metric.DeviceInfo, error) {
	var errSummary []string
	hostname, err := metric.GetHostname()
	if err != nil {
		errSummary = append(errSummary, "hostname: "+err.Error())
	}
	username, err := metric.GetCurrentUser()
	if err != nil {
		errSummary = append(errSummary, "username: "+err.Error())
	}
	osDisplay, err := metric.GetExactOSName()
	if err != nil {
		errSummary = append(errSummary, "os_name: "+err.Error())
	}
	domain, err := metric.GetWmiDomain()
	if err != nil {
		errSummary = append(errSummary, "domain: "+err.Error())
	}
	fw, err := firewall.GetFirewallStatus()
	fwStatus := false
	if err != nil {
		errSummary = append(errSummary, "firewall: "+err.Error())
	} else if fw != nil {
		fwStatus = fw.Enabled
	}

	def, err := GetDefenderStatus()

	defStatus := false

	tpmStatus := false

	if err != nil {
		errSummary = append(errSummary, "defender: "+err.Error())
	} else if def != nil {
		defStatus = def.AntivirusEnabled &&
			def.RealTimeProtection &&
			def.DefenderServiceState == "Running"

		tpmStatus = def.TamperProtection
	}
	data := metric.DeviceInfo{
		Hostname:        hostname,
		Domain:          domain,
		User:            username,
		OS:              osDisplay,
		AgentVersion:    "v1.0.1",
		WindowsDefender: defStatus,
		Firewall:        fwStatus,
		TPM:             tpmStatus,
		BitLocker:       false,
	}
	if len(errSummary) > 0 {
		return data, fmt.Errorf("device info metrics completed with errors: [%s]", strings.Join(errSummary, "; "))
	}
	return data, nil
}

func (a *App) DashboardUpdate() metric.DeviceInfo {
	data, err := GetSystemDeviceInfo()
	if err != nil {
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
	}
	return data
}
