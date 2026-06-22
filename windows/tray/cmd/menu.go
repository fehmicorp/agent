package main

import (
	"github.com/fehmicorp/win_tray/internal/config"
	"github.com/getlantern/systray"
)

var (
	Items []systray.MenuItem
)

func populateMenu(cfg *config.Tray) {
	for _, t := range cfg.Menu {
		if t.Separator == true {
			systray.AddSeparator()
		} else if t.Visible == true && t.Enabled == true {
			Items := systray.AddMenuItem(t.Title, t.Tooltip)
		} else if t.Visible == true && t.Enabled == false {

		}
	}
}

func onClick(item *systray.MenuItem, Id int) {
	for {
		<-item.ClickedCh
		handleTrayClick(Id)
	}
}

func handleTrayClick(Id int) {
	ctx := con
}
