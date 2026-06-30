package firewall

import (
	"encoding/json"
	"os/exec"
	"strings"

	"golang.org/x/sys/windows"
)

type FirewallStatus struct {
	Enabled         bool   `json:"enabled"`
	Status          string `json:"status"`
	DomainProfile   bool   `json:"domain"`
	PrivateProfile  bool   `json:"private"`
	PublicProfile   bool   `json:"public"`
	InboundAction   string `json:"inbound_action"`
	FirewallService string `json:"firewall_service"`
}

type firewallProfile struct {
	Name                 string      `json:"Name"`
	Enabled              interface{} `json:"Enabled"`
	DefaultInboundAction interface{} `json:"DefaultInboundAction"`
}

func parseBool(v interface{}) bool {
	switch x := v.(type) {
	case bool:
		return x
	case float64:
		return x == 1
	case int:
		return x == 1
	case string:
		switch strings.ToLower(strings.TrimSpace(x)) {
		case "true", "1", "enabled":
			return true
		}
	}
	return false
}

func parseAction(v interface{}) string {
	switch x := v.(type) {

	case string:
		switch strings.ToLower(strings.TrimSpace(x)) {
		case "allow":
			return "Allow"
		case "block":
			return "Block"
		}
		return x

	case float64:
		switch int(x) {
		case 1:
			return "Allow"
		case 2:
			return "Block"
		}

	case int:
		switch x {
		case 1:
			return "Allow"
		case 2:
			return "Block"
		}
	}

	return "Unknown"
}

func GetFirewallStatus() (*FirewallStatus, error) {

	result := &FirewallStatus{
		Status:          "Disabled",
		FirewallService: "Stopped",
		InboundAction:   "Unknown",
	}

	cmd := exec.Command(
		"powershell.exe",
		"-NoLogo",
		"-NoProfile",
		"-NonInteractive",
		"-ExecutionPolicy", "Bypass",
		"-Command",
		`Get-NetFirewallProfile |
		Select-Object Name,Enabled,DefaultInboundAction |
		ConvertTo-Json -Compress`,
	)

	// Run completely hidden
	cmd.SysProcAttr = &windows.SysProcAttr{
		HideWindow:    true,
		CreationFlags: windows.CREATE_NO_WINDOW,
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return result, err
	}

	var profiles []firewallProfile

	// PowerShell returns an object instead of an array when only one item exists
	if err := json.Unmarshal(out, &profiles); err != nil {
		var single firewallProfile
		if err2 := json.Unmarshal(out, &single); err2 != nil {
			return result, err
		}
		profiles = []firewallProfile{single}
	}

	result.FirewallService = "Running"

	for _, p := range profiles {

		enabled := parseBool(p.Enabled)

		if enabled {
			result.Enabled = true
			result.Status = "Enabled"
		}

		switch strings.TrimSpace(p.Name) {

		case "Domain":
			result.DomainProfile = enabled

		case "Private":
			result.PrivateProfile = enabled

		case "Public":
			result.PublicProfile = enabled
		}

		action := parseAction(p.DefaultInboundAction)
		if action != "Unknown" {
			result.InboundAction = action
		}
	}

	return result, nil
}
