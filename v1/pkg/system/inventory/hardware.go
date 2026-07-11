package inventory

import (
	"fmt"
	"runtime"

	"github.com/StackExchange/wmi"
)

func collectHardware(info *HardwareInfo) error {

	switch runtime.GOOS {

	case "windows":
		return collectWindowsHardware(info)

	case "linux":
		return collectLinuxHardware(info)

	case "darwin":
		return collectDarwinHardware(info)

	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// -----------------------------------------------------------------------------
// Windows
// -----------------------------------------------------------------------------

type winBIOSInfo struct {
	Manufacturer      string
	SMBIOSBIOSVersion string
}

type winBaseBoardInfo struct {
	Product string
}

type winSystemEnclosure struct {
	ChassisTypes []uint16
}

type winBattery struct {
	Name string
}

func collectWindowsHardware(info *HardwareInfo) error {

	// -------------------------------------------------------------------------
	// BIOS
	// -------------------------------------------------------------------------

	var bios []winBIOSInfo

	if err := wmi.Query(`
		SELECT
			Manufacturer,
			SMBIOSBIOSVersion
		FROM Win32_BIOS
	`, &bios); err == nil && len(bios) > 0 {

		info.BIOSVendor = bios[0].Manufacturer
		info.BIOSVersion = bios[0].SMBIOSBIOSVersion
	}

	// -------------------------------------------------------------------------
	// Motherboard
	// -------------------------------------------------------------------------

	var boards []winBaseBoardInfo

	if err := wmi.Query(`
		SELECT Product
		FROM Win32_BaseBoard
	`, &boards); err == nil && len(boards) > 0 {

		info.Motherboard = boards[0].Product
	}

	// -------------------------------------------------------------------------
	// Chassis
	// -------------------------------------------------------------------------

	var chassis []winSystemEnclosure

	if err := wmi.Query(`
		SELECT ChassisTypes
		FROM Win32_SystemEnclosure
	`, &chassis); err == nil && len(chassis) > 0 {

		info.Chassis = chassisType(chassis[0].ChassisTypes)
	}

	// -------------------------------------------------------------------------
	// Battery
	// -------------------------------------------------------------------------

	var batteries []winBattery

	if err := wmi.Query(`
		SELECT Name
		FROM Win32_Battery
	`, &batteries); err == nil {

		info.Battery = len(batteries) > 0
	}

	return nil
}

// -----------------------------------------------------------------------------
// Linux
// -----------------------------------------------------------------------------

func collectLinuxHardware(info *HardwareInfo) error {

	// TODO:
	// Read from:
	// /sys/class/dmi/id
	// /sys/class/power_supply

	return nil
}

// -----------------------------------------------------------------------------
// macOS
// -----------------------------------------------------------------------------

func collectDarwinHardware(info *HardwareInfo) error {

	// TODO:
	// Use system_profiler
	// Use ioreg

	return nil
}

// -----------------------------------------------------------------------------
// Helpers
// -----------------------------------------------------------------------------

func chassisType(types []uint16) string {

	if len(types) == 0 {
		return ""
	}

	switch types[0] {

	case 3:
		return "Desktop"

	case 4:
		return "Low Profile Desktop"

	case 5:
		return "Pizza Box"

	case 6:
		return "Mini Tower"

	case 7:
		return "Tower"

	case 8:
		return "Portable"

	case 9:
		return "Laptop"

	case 10:
		return "Notebook"

	case 11:
		return "Handheld"

	case 12:
		return "Docking Station"

	case 13:
		return "All-in-One"

	case 14:
		return "Sub Notebook"

	case 15:
		return "Space Saving"

	case 16:
		return "Lunch Box"

	case 23:
		return "Rack Mount"

	case 30:
		return "Tablet"

	case 31:
		return "Convertible"

	case 32:
		return "Detachable"

	default:
		return fmt.Sprintf("Type %d", types[0])
	}
}
