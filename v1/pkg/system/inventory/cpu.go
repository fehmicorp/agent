package inventory

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/StackExchange/wmi"
)

func collectCPU(info *CPUInfo) error {

	switch runtime.GOOS {

	case "windows":
		return collectWindowsCPU(info)

	case "linux":
		return collectLinuxCPU(info)

	case "darwin":
		return collectDarwinCPU(info)

	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// -----------------------------------------------------------------------------
// Windows
// -----------------------------------------------------------------------------

type winProcessor struct {
	Name                      string
	Manufacturer              string
	NumberOfCores             uint32
	NumberOfLogicalProcessors uint32
	MaxClockSpeed             uint32
	CurrentClockSpeed         uint32
}

func collectWindowsCPU(info *CPUInfo) error {

	var cpus []winProcessor

	err := wmi.Query(`
		SELECT
			Name,
			Manufacturer,
			NumberOfCores,
			NumberOfLogicalProcessors,
			MaxClockSpeed,
			CurrentClockSpeed
		FROM Win32_Processor
	`, &cpus)

	if err != nil {
		return err
	}

	if len(cpus) == 0 {
		return nil
	}

	// Aggregate values for multi-CPU systems.
	for _, cpu := range cpus {

		if info.Model == "" {
			info.Model = cpu.Name
		}

		if info.Manufacturer == "" {
			info.Manufacturer = cpu.Manufacturer
		}

		info.Cores += int(cpu.NumberOfCores)
		info.LogicalCores += int(cpu.NumberOfLogicalProcessors)

		if int(cpu.MaxClockSpeed) > info.MaxClockMHz {
			info.MaxClockMHz = int(cpu.MaxClockSpeed)
		}

		if int(cpu.CurrentClockSpeed) > info.CurrentClockMHz {
			info.CurrentClockMHz = int(cpu.CurrentClockSpeed)
		}
	}

	return nil
}

// -----------------------------------------------------------------------------
// Linux
// -----------------------------------------------------------------------------

func collectLinuxCPU(info *CPUInfo) error {

	// Model
	if out, err := exec.Command("sh", "-c", "grep 'model name' /proc/cpuinfo | head -1 | cut -d ':' -f2").Output(); err == nil {
		info.Model = strings.TrimSpace(string(out))
	}

	// Vendor
	if out, err := exec.Command("sh", "-c", "grep 'vendor_id' /proc/cpuinfo | head -1 | cut -d ':' -f2").Output(); err == nil {
		info.Manufacturer = strings.TrimSpace(string(out))
	}

	// Physical cores
	if out, err := exec.Command("nproc", "--all").Output(); err == nil {
		if v, err := strconv.Atoi(strings.TrimSpace(string(out))); err == nil {
			info.LogicalCores = v
		}
	}

	// Cores per socket (best effort)
	if out, err := exec.Command("sh", "-c", "lscpu | grep '^Core(s) per socket:' | awk '{print $4}'").Output(); err == nil {
		if v, err := strconv.Atoi(strings.TrimSpace(string(out))); err == nil {
			info.Cores = v
		}
	}

	// Max MHz
	if out, err := exec.Command("sh", "-c", "lscpu | grep '^CPU max MHz:' | awk '{print int($4)}'").Output(); err == nil {
		if v, err := strconv.Atoi(strings.TrimSpace(string(out))); err == nil {
			info.MaxClockMHz = v
		}
	}

	// Current MHz
	if out, err := exec.Command("sh", "-c", "lscpu | grep '^CPU MHz:' | awk '{print int($3)}'").Output(); err == nil {
		if v, err := strconv.Atoi(strings.TrimSpace(string(out))); err == nil {
			info.CurrentClockMHz = v
		}
	}

	// Fallback
	if info.Cores == 0 {
		info.Cores = runtime.NumCPU()
	}
	if info.LogicalCores == 0 {
		info.LogicalCores = runtime.NumCPU()
	}

	return nil
}

// -----------------------------------------------------------------------------
// macOS
// -----------------------------------------------------------------------------

func collectDarwinCPU(info *CPUInfo) error {

	read := func(key string) string {
		out, err := exec.Command("sysctl", "-n", key).Output()
		if err != nil {
			return ""
		}
		return strings.TrimSpace(string(out))
	}

	info.Model = read("machdep.cpu.brand_string")

	// Apple Silicon returns "Apple"
	if v := read("machdep.cpu.vendor"); v != "" {
		info.Manufacturer = v
	} else {
		info.Manufacturer = "Apple"
	}

	if v, err := strconv.Atoi(read("hw.physicalcpu")); err == nil {
		info.Cores = v
	}

	if v, err := strconv.Atoi(read("hw.logicalcpu")); err == nil {
		info.LogicalCores = v
	}

	if v, err := strconv.Atoi(read("hw.cpufrequency_max")); err == nil {
		info.MaxClockMHz = v / 1000000
	}

	if v, err := strconv.Atoi(read("hw.cpufrequency")); err == nil {
		info.CurrentClockMHz = v / 1000000
	}

	if info.Cores == 0 {
		info.Cores = runtime.NumCPU()
	}

	if info.LogicalCores == 0 {
		info.LogicalCores = runtime.NumCPU()
	}

	return nil
}
