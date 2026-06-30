package metric

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"unicode"

	"github.com/StackExchange/wmi"
	"github.com/fehmicorp/agent/windows/debug/logger"
	"github.com/shirou/gopsutil/host"
)

type Win32_ComputerSystem struct {
	Domain       string
	PartOfDomain bool
}

type DeviceInfo struct {
	Hostname        string `json:"hostname"`
	Domain          string `json:"domain"`
	User            string `json:"user"`
	OS              string `json:"os"`
	AgentVersion    string `json:"agentVersion"`
	WindowsDefender string `json:"windowsDefender"`
	Firewall        string `json:"firewall"`
	TPM             string `json:"tpm"`
	BitLocker       string `json:"bitLocker"`
}

func (a *App) GetSystemDeviceInfo() DeviceInfo {
	hostname, _ := os.Hostname()
	currentUser, _ := user.Current()
	osDisplay := getExactOSName()
	username := "Administrator"
	domain := formatToTitleCase(getWmiDomain())
	// defender := getDefender()
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

	logger.Logger.Log("INFO", "Device info layout requested and built successfully")
	logger.Logger.Log("SECURITY", fmt.Sprintf("Defender Overall: %s | Realtime: %s | Tamper: %s", defender.Status, defender.RealTimeProtection, defender.TamperProtection))
	logger.Logger.Log("SECURITY", fmt.Sprintf("Firewall Overall: %s | Private: %s | Public: %s", firewall.Status, firewall.PrivateProfile, firewall.PublicProfile))
	return DeviceInfo{
		Hostname:     hostname,
		Domain:       domain,
		User:         username,
		OS:           osDisplay,
		AgentVersion: "v1.0.1",
		// WindowsDefender: defender.Status,
		Firewall:  firewall.Status,
		TPM:       "Enabled",
		BitLocker: "Disabled",
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
