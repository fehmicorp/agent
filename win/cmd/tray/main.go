package mtray

import (
	"log"
	"win/internal/config"
	"win/internal/tray"
)

func RunTray() {
	config.InitialLoad()
	if err := tray.Init(); err != nil {
		log.Fatal(err)
	}
	tray.Run()
}
