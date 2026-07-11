package tray

import (
	"github.com/getlantern/systray"
)

func RunTray() {
	LoadTrayConfig()
	systray.Run(onReady, onExit)
}
