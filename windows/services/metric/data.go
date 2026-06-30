package metric

import (
	"strings"
	"unicode"

	"github.com/StackExchange/wmi"
	"github.com/shirou/gopsutil/host"
)

type Win32_ComputerSystem struct {
	Domain       string
	PartOfDomain bool
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
