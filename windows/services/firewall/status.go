package firewall

import (
	"encoding/json"
	"os/exec"
	"strings"
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
		switch strings.ToLower(x) {
		case "true", "1", "enabled":
			return true
		}
	}
	return false
}

func parseAction(v interface{}) string {
	switch x := v.(type) {
	case string:
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
		"powershell",
		"-NoProfile",
		"-ExecutionPolicy", "Bypass",
		"-Command",
		`Get-NetFirewallProfile | Select Name,Enabled,DefaultInboundAction | ConvertTo-Json`,
	)
	out, err := cmd.Output()
	if err != nil {
		return result, err
	}
	var profiles []firewallProfile
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
		switch p.Name {
		case "Domain":
			result.DomainProfile = enabled
		case "Private":
			result.PrivateProfile = enabled

		case "Public":
			result.PublicProfile = enabled
		}
		if p.DefaultInboundAction != "" {
			result.InboundAction = parseAction(p.DefaultInboundAction)
		}
	}
	return result, nil
}
