package defender

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

type TPMStatus struct {
	Present   bool `json:"present"`
	Ready     bool `json:"ready"`
	Enabled   bool `json:"enabled"`
	Activated bool `json:"activated"`
	Owned     bool `json:"owned"`
}

type tpmResponse struct {
	TpmPresent   bool `json:"TpmPresent"`
	TpmReady     bool `json:"TpmReady"`
	TpmEnabled   bool `json:"TpmEnabled"`
	TpmActivated bool `json:"TpmActivated"`
	TpmOwned     bool `json:"TpmOwned"`
}

func GetTPMStatus() (*TPMStatus, error) {

	result := &TPMStatus{
		Present:   false,
		Ready:     false,
		Enabled:   false,
		Activated: false,
		Owned:     false,
	}

	script := "Get-Tpm | Select-Object TpmPresent,TpmReady,TpmEnabled,TpmActivated,TpmOwned | ConvertTo-Json -Compress"

	cmd := exec.Command(
		"powershell.exe",
		"-NoLogo",
		"-NoProfile",
		"-NonInteractive",
		"-ExecutionPolicy", "Bypass",
		"-Command",
		script,
	)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: windows.CREATE_NO_WINDOW,
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, string(out))
	}

	// Optional for debugging
	fmt.Println(string(out))

	var ps tpmResponse
	if err := json.Unmarshal(out, &ps); err != nil {
		return nil, fmt.Errorf("failed to parse TPM JSON: %w", err)
	}

	result.Present = ps.TpmPresent
	result.Ready = ps.TpmReady
	result.Enabled = ps.TpmEnabled
	result.Activated = ps.TpmActivated
	result.Owned = ps.TpmOwned

	return result, nil
}
