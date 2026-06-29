package main

import (
	"fmt"
	"os"
	"os/user"
	goruntime "runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) startMetricReportingLoop() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	prevNet, _ := net.IOCounters(true)
	lastSample := time.Now()

	for {

		select {

		case <-a.ctx.Done():
			return

		case <-ticker.C:

			// ---------------- CPU ----------------

			cpuPercent, _ := cpu.Percent(0, false)
			cpuStr := "0%"

			if len(cpuPercent) > 0 {
				cpuStr = fmt.Sprintf("%.0f%%", cpuPercent[0])
			}

			// ---------------- RAM ----------------

			ramStr := "0%"

			if vm, err := mem.VirtualMemory(); err == nil {
				ramStr = fmt.Sprintf("%.0f%%", vm.UsedPercent)
			}

			// ---------------- Disk ----------------

			diskStr := "0%"

			d, err := disk.Usage("/")
			if err != nil {
				d, _ = disk.Usage("C:")
			}

			if d != nil {
				diskStr = fmt.Sprintf("%.0f%%", d.UsedPercent)
			}

			// ---------------- Network ----------------

			netStr := "↓ 0B | ↑ 0B"

			currNet, err := net.IOCounters(true)

			if err == nil && len(currNet) > 0 && len(prevNet) > 0 {

				now := time.Now()
				elapsed := now.Sub(lastSample).Seconds()

				if elapsed <= 0 {
					elapsed = 1
				}

				var rxDiff uint64
				var txDiff uint64

				if currNet[0].BytesRecv >= prevNet[0].BytesRecv {
					rxDiff = currNet[0].BytesRecv - prevNet[0].BytesRecv
				}

				if currNet[0].BytesSent >= prevNet[0].BytesSent {
					txDiff = currNet[0].BytesSent - prevNet[0].BytesSent
				}

				rxSpeed := float64(rxDiff) / elapsed
				txSpeed := float64(txDiff) / elapsed

				// Reject impossible values (>10 Gbps ≈ 1.25 GB/s)
				if rxSpeed > 1.25*1024*1024*1024 {
					rxSpeed = 0
				}

				if txSpeed > 1.25*1024*1024*1024 {
					txSpeed = 0
				}

				netStr = fmt.Sprintf(
					"↓ %s | ↑ %s",
					formatSpeed(rxSpeed),
					formatSpeed(txSpeed),
				)

				prevNet = currNet
				lastSample = now
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

func formatSpeed(bytes float64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.2fGB", bytes/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2fMB", bytes/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.0fKB", bytes/KB)
	default:
		return fmt.Sprintf("%.0fB", bytes)
	}
}

func (a *App) GetSystemDeviceInfo() DeviceInfo {
	hostname, _ := os.Hostname()
	currentUser, _ := user.Current()

	osType := goruntime.GOOS
	osDisplay := "Linux / macOS"
	if osType == "windows" {
		osDisplay = "Windows 11 Pro"
	}

	username := "Administrator"
	if currentUser != nil {
		username = currentUser.Username
	}

	return DeviceInfo{
		Hostname:        hostname,
		Domain:          "Workgroup",
		User:            username,
		OS:              osDisplay,
		AgentVersion:    "v1.0.1",
		WindowsDefender: "Enabled",
		Firewall:        "Enabled",
		TPM:             "Enabled",
		BitLocker:       "Disabled",
	}
}
