package runas

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

func IsAdmin() bool {
	token := windows.GetCurrentProcessToken()
	adminSid, err := windows.CreateWellKnownSid(windows.WinBuiltinAdministratorsSid)
	if err != nil {
		return false
	}
	member, err := token.IsMember(adminSid)
	if err != nil {
		return false
	}
	return member
}

func RequestElevation() error {
	if IsAdmin() {
		return nil // Already elevated
	}

	verb := "runas"
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to detect running executable location: %w", err)
	}

	// Forward all incoming flags or parameters seamlessly to the elevated instance
	args := strings.Join(os.Args[1:], " ")

	verbPtr, err := syscall.UTF16PtrFromString(verb)
	if err != nil {
		return err
	}
	exePtr, err := syscall.UTF16PtrFromString(exe)
	if err != nil {
		return err
	}
	cwdPtr, err := syscall.UTF16PtrFromString(os.Getenv("CD"))
	if err != nil {
		return err
	}
	argsPtr, err := syscall.UTF16PtrFromString(args)
	if err != nil {
		return err
	}

	var showCmd int32 = windows.SW_NORMAL

	// Fill standard Shell Execute Struct parameters
	sei := &windows.ShellExecuteInfo{
		Size:       uint32(windows.SizeofShellExecuteInfo),
		Mask:       windows.SEE_MASK_NOASYNC,
		Wnd:        0,
		Verb:       verbPtr,
		File:       exePtr,
		Parameters: argsPtr,
		Directory:  cwdPtr,
		Show:       showCmd,
		HInstApp:   0,
	}

	// Trigger execution through Win32 subsystem API
	err = windows.ShellExecuteEx(sei)
	if err != nil {
		return fmt.Errorf("elevation prompt rejected or failed: %w", err)
	}

	// Exit the unprivileged instance as the new process context spawns successfully
	os.Exit(0)
	return nil
}
