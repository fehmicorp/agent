package tray

import (
	_ "embed"
	"log"
	"win/internal/app"
	"win/internal/assets"
	"win/internal/config"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var tryCfg *config.Tray

func onReady() {
	if tryCfg == nil {
		panic("tray config not loaded")
	}
	systray.SetIcon(assets.Favicon)
	systray.SetTitle(tryCfg.Title)
	systray.SetTooltip(tryCfg.Tooltip)
	go populateMenu()
}

func populateMenu() {
	for _, t := range tryCfg.Menu {
		if t.Separator == true {
			systray.AddSeparator()
		} else if t.Visible == true && t.Enabled == true {
			item := systray.AddMenuItem(t.Title, t.Tooltip)
			go func(item *systray.MenuItem, id int) {
				for {
					<-item.ClickedCh
					handleTrayClick(t.FuncId)
				}
			}(item, t.FuncId)
		}
	}
}

func handleTrayClick(id int) {
	ctx := config.GetContext(id)
	if ctx != nil {
		runtime.WindowShow(ctx)
		runtime.WindowCenter(ctx)
	} else {
		go func(id int) {
			log.Printf("Launching new window for ID: %d", id)
			app.RunApp(id)
		}(id)
	}
}
