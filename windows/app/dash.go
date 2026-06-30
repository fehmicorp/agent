package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fehmicorp/agent/windows/debug/logger"
	log "github.com/fehmicorp/agent/windows/debug/logger"
	"github.com/fehmicorp/agent/windows/services/metric"
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

func (a *App) DashboardUpdate() metric.DeviceInfo {
	data, err := metric.GetSystemDeviceInfo()
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
