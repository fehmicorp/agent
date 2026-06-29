package main

import (
	"fmt"
	"os/exec"

	"golang.org/x/sys/windows"
)

func setFirewallEnabled() error {
	return setFirewall(true)
}

func setFirewallDisabled() error {
	return setFirewall(false)
}

func setFirewall(enable bool) error {

	state := "False"
	if enable {
		state = "True"
	}

	cmd := exec.Command(
		"powershell",
		"-NoProfile",
		"-ExecutionPolicy", "Bypass",
		"-Command",
		fmt.Sprintf(
			`Set-NetFirewallProfile -Profile Domain,Private,Public -Enabled %s`,
			state,
		),
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, string(out))
	}

	return nil
}

func IsAdmin() bool {
	token := windows.GetCurrentProcessToken()

	adminSid, _ := windows.CreateWellKnownSid(
		windows.WinBuiltinAdministratorsSid,
	)

	member, err := token.IsMember(adminSid)
	if err != nil {
		return false
	}

	return member
}

func setFirewallProfile(profile string, enable bool) error {
	state := "False"
	if enable {
		state = "True"
	}
	cmd := exec.Command(
		"powershell",
		"-NoProfile",
		"-ExecutionPolicy", "Bypass",
		"-Command",
		fmt.Sprintf(
			`Set-NetFirewallProfile -Profile %s -Enabled %s`,
			profile,
			state,
		),
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, string(out))
	}
	return nil
}
