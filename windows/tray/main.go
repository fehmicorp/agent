package main

import (
	"github.com/fehmicorp/agent/windows/config/load"
	"github.com/getlantern/systray"
)

func main() {
	load.TrayConf()
	// if err := tray.Init(); err != nil {
	// 	log.Fatal(err)
	// }
	systray.Run(onReady, onReady)
}
