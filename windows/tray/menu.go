package main

import (
	"log"

	"github.com/fehmicorp/agent/windows/config/registry"
	types "github.com/fehmicorp/agent/windows/config/types"
	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func populateMenu(cfg *types.Tray) {
	for _, t := range cfg.Menu {
		if t.Separator == true {
			systray.AddSeparator()
		} else if t.Visible && t.Enabled {
			items := systray.AddMenuItem(t.Title, t.Tooltip)
			systray.SetTooltip(t.Tooltip)
			go onClick(items, t.ID)
		} else if t.Visible && !t.Enabled {
			items := systray.AddMenuItem(t.Title, "")
			items.Disable()
		}
	}
}

func onClick(items *systray.MenuItem, id int) {
	for range items.ClickedCh {
		handleTrayClick(id)
	}
}

func handleTrayClick(Id int) {
	if Id == 9999 {
		log.Println("Exit menu option clicked. Shutting down tray...")
		systray.Quit()
		return
	}
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
