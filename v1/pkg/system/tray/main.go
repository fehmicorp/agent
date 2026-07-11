package tray

import (
	"github.com/getlantern/systray"
)

func main() {
	LoadTrayConfig()
	systray.Run(onReady, onExit)
}
