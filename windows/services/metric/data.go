package metric

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"unicode"

	"github.com/StackExchange/wmi"
	"github.com/fehmicorp/agent/windows/services/bitlocker"
	"github.com/fehmicorp/agent/windows/services/defender"
	"github.com/fehmicorp/agent/windows/services/firewall"
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
	WindowsDefender bool   `json:"windowsDefender"`
	Firewall        bool   `json:"firewall"`
	TPM             bool   `json:"tpm"`
	BitLocker       bool   `json:"bitLocker"`
}

func GetHostname() (string, error) {
	data, err := os.Hostname()
	return data, err
}

func GetCurrentUser() (string, error) {
	currentUser := ""
	data, err := user.Current()
	if data != nil {
		rawUsername := data.Username
		if strings.Contains(rawUsername, "\\") {
			parts := strings.Split(rawUsername, "\\")
			currentUser = parts[len(parts)-1]
		} else {
			currentUser = rawUsername
		}
	}
	return currentUser, err
}

func GetSystemDeviceInfo() (DeviceInfo, error) {
	var errSummary []string
	hostname, err := GetHostname()
	if err != nil {
		errSummary = append(errSummary, "hostname: "+err.Error())
	}
	username, err := GetCurrentUser()
	if err != nil {
		errSummary = append(errSummary, "username: "+err.Error())
	}
	osDisplay, err := GetExactOSName()
	if err != nil {
		errSummary = append(errSummary, "os_name: "+err.Error())
	}
	domain, err := GetWmiDomain()
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

	def, err := defender.GetDefenderStatus()
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
	bitlockerStatus := false
	blVolumes, err := bitlocker.GetBitLockerStatus()
	if err != nil {
		errSummary = append(errSummary, "bitlocker: "+err.Error())
	} else {
		for _, vol := range blVolumes {
			if strings.EqualFold(vol.Drive, "C:") &&
				vol.Encrypted &&
				vol.ProtectionEnabled {
				bitlockerStatus = true
				break
			}
		}
	}
	data := DeviceInfo{
		Hostname:        hostname,
		Domain:          domain,
		User:            username,
		OS:              osDisplay,
		AgentVersion:    "",
		WindowsDefender: defStatus,
		Firewall:        fwStatus,
		TPM:             tpmStatus,
		BitLocker:       bitlockerStatus,
	}
	if len(errSummary) > 0 {
		return data, fmt.Errorf("device info metrics completed with errors: [%s]", strings.Join(errSummary, "; "))
	}
	return data, nil
}

func GetWmiDomain() (string, error) {
	var systemInfo []Win32_ComputerSystem
	query := "SELECT Domain, PartOfDomain FROM Win32_ComputerSystem"
	if err := wmi.Query(query, &systemInfo); err != nil || len(systemInfo) == 0 {
		return formatToTitleCase("Workgroup"), err
	}
	sys := systemInfo[0]
	if sys.Domain != "" {
		return formatToTitleCase(sys.Domain), nil
	}
	return formatToTitleCase("Workgroup"), nil
}

func GetExactOSName() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "Unknown OS", err
	}
	if info.Platform != "" {
		return info.Platform, nil
	}
	return info.OS, nil
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
