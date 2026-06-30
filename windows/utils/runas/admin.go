package runas

import (
	"fmt"
	"os"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	shell32           = windows.NewLazySystemDLL("shell32.dll")
	procShellExecuteW = shell32.NewProc("ShellExecuteW")
)

func IsAdmin() bool {
	token := windows.GetCurrentProcessToken()

	adminSID, err := windows.CreateWellKnownSid(
		windows.WinBuiltinAdministratorsSid,
	)
	if err != nil {
		return false
	}

	member, err := token.IsMember(adminSID)
	if err != nil {
		return false
	}

	return member
}

func RequestElevation() error {

	if IsAdmin() {
		return nil
	}

	exe, err := os.Executable()
	if err != nil {
		return err
	}

	args := strings.Join(os.Args[1:], " ")

	verb, _ := windows.UTF16PtrFromString("runas")
	file, _ := windows.UTF16PtrFromString(exe)
	params, _ := windows.UTF16PtrFromString(args)

	dir, _ := os.Getwd()
	workdir, _ := windows.UTF16PtrFromString(dir)

	const SW_NORMAL = 1

	r1, _, err := procShellExecuteW.Call(
		0,
		uintptr(unsafe.Pointer(verb)),
		uintptr(unsafe.Pointer(file)),
		uintptr(unsafe.Pointer(params)),
		uintptr(unsafe.Pointer(workdir)),
		uintptr(SW_NORMAL),
	)

	// Per Microsoft docs:
	// Return value > 32 means success.
	if r1 <= 32 {
		return fmt.Errorf("ShellExecuteW failed: %v (code=%d)", err, r1)
	}

	os.Exit(0)
	return nil
}
