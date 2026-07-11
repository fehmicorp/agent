package inventory

import (
	"fmt"
	"runtime"
	"time"

	"github.com/StackExchange/wmi"
)

func collectOS(info *OSInfo) error {

	switch runtime.GOOS {

	case "windows":
		return collectWindowsOS(info)

	case "linux":
		return collectLinuxOS(info)

	case "darwin":
		return collectDarwinOS(info)

	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// -----------------------------------------------------------------------------
// Windows
// -----------------------------------------------------------------------------

type winOperatingSystem struct {
	Caption        string
	Version        string
	BuildNumber    string
	InstallDate    string
	LastBootUpTime string
}

func collectWindowsOS(info *OSInfo) error {

	info.Platform = runtime.GOOS
	info.Architecture = runtime.GOARCH

	var osInfo []winOperatingSystem

	if err := wmi.Query(`
		SELECT
			Caption,
			Version,
			BuildNumber,
			InstallDate,
			LastBootUpTime
		FROM Win32_OperatingSystem
	`, &osInfo); err != nil {
		return err
	}

	if len(osInfo) == 0 {
		return nil
	}

	os := osInfo[0]

	info.Name = os.Caption
	info.Version = os.Version
	info.Build = os.BuildNumber
	info.Kernel = os.Version

	info.InstallDate = parseWMITime(os.InstallDate)
	info.BootTime = parseWMITime(os.LastBootUpTime)

	if boot, err := parseWMITimeValue(os.LastBootUpTime); err == nil {
		info.Uptime = uint64(time.Since(boot).Seconds())
	}

	return nil
}

// -----------------------------------------------------------------------------
// Linux
// -----------------------------------------------------------------------------

func collectLinuxOS(info *OSInfo) error {

	info.Platform = runtime.GOOS
	info.Architecture = runtime.GOARCH

	return nil
}

// -----------------------------------------------------------------------------
// macOS
// -----------------------------------------------------------------------------

func collectDarwinOS(info *OSInfo) error {

	info.Platform = runtime.GOOS
	info.Architecture = runtime.GOARCH

	return nil
}

// -----------------------------------------------------------------------------
// Helpers
// -----------------------------------------------------------------------------

func parseWMITime(v string) string {

	t, err := parseWMITimeValue(v)
	if err != nil {
		return ""
	}

	return t.Format(time.RFC3339)
}

func parseWMITimeValue(v string) (time.Time, error) {

	// WMI format:
	// 20260711132430.000000+330

	if len(v) < 14 {
		return time.Time{}, fmt.Errorf("invalid WMI datetime")
	}

	return time.Parse(
		"20060102150405",
		v[:14],
	)
}
