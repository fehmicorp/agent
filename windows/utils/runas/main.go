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
		return nil
	}

	verb := "runas"
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to detect running executable location: %w", err)
	}

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

	// Use windows namespace for the API types
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

	// Correct native Windows wrapper call
	err = windows.ShellExecuteEx(sei)
	if err != nil {
		return fmt.Errorf("elevation prompt rejected or failed: %w", err)
	}

	os.Exit(0)
	return nil
}
