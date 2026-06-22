package main

import (
	"log"
	"runtime"

	"github.com/fehmicorp/agent/windows/config/registry"
	"github.com/fehmicorp/win_tray/internal/config"
	"github.com/getlantern/systray"
)

func populateMenu(cfg *config.Tray) {
	for _, t := range cfg.Menu {
		if t.Separator == true {
			systray.AddSeparator()
		} else if t.Visible == true && t.Enabled == true {
			Items := systray.AddMenuItem(t.Title, t.Tooltip)
			onClick(Items, t.ID)
		} else if t.Visible == true && t.Enabled == false {

		}
	}
}

func onClick(Items *systray.MenuItem, Id int) {
	for {
		<-Items.ClickedCh
		handleTrayClick(Id)
	}
}

func handleTrayClick(Id int) {
	ctx := registry.GetContext(Id)
	if ctx != nil {
		runtime.WindowShow(ctx)
		runtime.WindowCenter(ctx)
	} else {
		go func(id int) {
			log.Printf("Launching new window for ID: %d", id)
		}(Id)
	}
}
