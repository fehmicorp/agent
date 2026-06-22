package tray

import (
	"github.com/getlantern/systray"
)

func Run() {
	systray.Run(onReady, onExit)
}
