package tray

import (
	"log"

	"github.com/fehmicorp/agent/windows/config/load"
	"github.com/getlantern/systray"
)

func main() {
	var err error
	cfg, err = load.TrayConf()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	systray.Run(onReady, onExit)
}
