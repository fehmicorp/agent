package tray

import (
	"log"
	"os/exec"
)

func syncNow() {
	log.Println("manual sync requested")
}

func requestSupport() {
	log.Println("support requested")
}

func restartAgent() {
	log.Println("agent restart requested")
}

func openLogs() {

	exec.Command(
		"explorer",
		`C:\ProgramData\Fehmi\logs`,
	).Start()
}

func onExit() {
	log.Println("tray closed")
}
