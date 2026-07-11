package inventory

import (
	"fmt"
	"net"
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/StackExchange/wmi"
)

func collectSystem(info *SystemInfo) error {
	switch runtime.GOOS {

	case "windows":
		return collectWindowsSystem(info)

	case "linux":
		return collectLinuxSystem(info)

	case "darwin":
		return collectDarwinSystem(info)

	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

type winComputerSystem struct {
	Name         string
	Domain       string
	Workgroup    string
	Manufacturer string
	Model        string
}

type winComputerSystemProduct struct {
	UUID string
}

type winBIOS struct {
	SerialNumber string
}

type winBaseBoard struct {
	SerialNumber string
}

type winOS struct {
	RegisteredUser string
}

func collectWindowsSystem(info *SystemInfo) error {

	host, _ := os.Hostname()
	info.Hostname = host

	if u, err := user.Current(); err == nil {
		username := u.Username
		if i := strings.LastIndex(username, "\\"); i >= 0 {
			username = username[i+1:]
		}
		info.User = username
		info.DisplayName = u.Name
	}

	// -------------------------------------------------------------------------
	// Computer System
	// -------------------------------------------------------------------------

	var systems []winComputerSystem

	if err := wmi.Query(`
		SELECT
			Name,
			Domain,
			Workgroup,
			Manufacturer,
			Model
		FROM Win32_ComputerSystem
	`, &systems); err == nil && len(systems) > 0 {

		info.Domain = systems[0].Domain
		info.Workgroup = systems[0].Workgroup
		info.Manufacturer = systems[0].Manufacturer
		info.Model = systems[0].Model
	}

	// -------------------------------------------------------------------------
	// Registered User
	// -------------------------------------------------------------------------

	var osInfo []winOS

	if err := wmi.Query(`
		SELECT RegisteredUser
		FROM Win32_OperatingSystem
	`, &osInfo); err == nil && len(osInfo) > 0 {

		if info.DisplayName == "" {
			info.DisplayName = osInfo[0].RegisteredUser
		}
	}

	// -------------------------------------------------------------------------
	// BIOS
	// -------------------------------------------------------------------------

	var bios []winBIOS

	if err := wmi.Query(`
		SELECT SerialNumber
		FROM Win32_BIOS
	`, &bios); err == nil && len(bios) > 0 {

		info.BIOSSerial = bios[0].SerialNumber
	}

	// -------------------------------------------------------------------------
	// Motherboard
	// -------------------------------------------------------------------------

	var boards []winBaseBoard

	if err := wmi.Query(`
		SELECT SerialNumber
		FROM Win32_BaseBoard
	`, &boards); err == nil && len(boards) > 0 {

		info.BoardSerial = boards[0].SerialNumber
	}

	// -------------------------------------------------------------------------
	// UUID
	// -------------------------------------------------------------------------

	var products []winComputerSystemProduct

	if err := wmi.Query(`
		SELECT UUID
		FROM Win32_ComputerSystemProduct
	`, &products); err == nil && len(products) > 0 {

		info.UUID = products[0].UUID
	}

	// -------------------------------------------------------------------------
	// System Serial Number
	// -------------------------------------------------------------------------

	info.SerialNumber = info.BIOSSerial

	// -------------------------------------------------------------------------
	// Primary MAC
	// -------------------------------------------------------------------------

	if ifaces, err := net.Interfaces(); err == nil {

		for _, iface := range ifaces {

			if iface.Flags&net.FlagUp == 0 {
				continue
			}

			if iface.Flags&net.FlagLoopback != 0 {
				continue
			}

			if iface.HardwareAddr.String() == "" {
				continue
			}

			info.PrimaryMAC = iface.HardwareAddr.String()
			break
		}
	}

	return nil
}

func collectLinuxSystem(info *SystemInfo) error {

	host, _ := os.Hostname()
	info.Hostname = host

	if u, err := user.Current(); err == nil {
		info.User = u.Username
		info.DisplayName = u.Name
	}

	return nil
}

func collectDarwinSystem(info *SystemInfo) error {

	host, _ := os.Hostname()
	info.Hostname = host

	if u, err := user.Current(); err == nil {
		info.User = u.Username
		info.DisplayName = u.Name
	}

	return nil
}
