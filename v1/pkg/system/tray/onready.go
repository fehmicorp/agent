package tray

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/fehmicorp/agent/v1/res/assets"
	"github.com/fehmicorp/agent/v1/res/logger"
	"github.com/getlantern/systray"
)

func onReady() {
	if cfg == nil {
		logger.Warn("System tray config not loaded")
	}
	icon, _ := assets.Read("fav.ico")
	systray.SetIcon(icon)
	systray.SetTitle(cfg.Title)
	systray.SetTooltip(fmt.Sprintf("%s\n%s", cfg.Tooltip, cfg.Version))
	populateMenu(cfg)
}

func populateMenu(cfg *Tray) {
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
		logger.Warn("Exit menu option clicked. Shutting down tray...")
		systray.Quit()
		return
	}
}

func LaunchApp(id int) error {
	exe := filepath.Join(
		".",
		"app.exe",
	)
	cmd := exec.Command(
		exe,
		"-window",
		fmt.Sprintf("%d", id),
	)
	return cmd.Start()
}
