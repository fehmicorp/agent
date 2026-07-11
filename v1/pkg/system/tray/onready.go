package tray

import (
	"fmt"
	"os/exec"

	"github.com/fehmicorp/agent/v1/cmd/appinfo"
	"github.com/fehmicorp/agent/v1/res/assets"
	"github.com/fehmicorp/agent/v1/res/logger"
	"github.com/getlantern/systray"
)

var appCmd *exec.Cmd

func onReady() {
	if cfg == nil {
		logger.Warn("System tray config not loaded")
	} else {
		icon, _ := assets.Read(cfg.Icon)
		systray.SetIcon(icon)
		systray.SetTitle(cfg.Title)
		systray.SetTooltip(fmt.Sprintf("%s\n%s", cfg.Tooltip, cfg.Version))
		populateMenu(cfg)
	}
}

func populateMenu(cfg *Tray) {
	for _, t := range cfg.Menu {
		if t.Separator == true {
			systray.AddSeparator()
			continue
		}
		if !t.Visible {
			continue
		}
		item := systray.AddMenuItem(t.Title, t.Tooltip)
		if t.Icon != "" {
			if icon, err := assets.Read("tray/" + t.Icon + ".ico"); err == nil {
				item.SetIcon(icon)
			} else {
				logger.Warn("Unable to load tray icon ", t.Icon, err)
			}
		}
		if !t.Enabled {
			item.Disable()
			continue
		}
		go onClick(item, t.FuncId)
	}
}

func onClick(items *systray.MenuItem, id int) {
	for range items.ClickedCh {
		handleTrayClick(id)
	}
}

func handleTrayClick(id int) {
	switch id {
	case 9999:
		logger.Warn("Exit menu option clicked. Shutting down tray...")
		systray.Quit()
		return

	case 1001:
		logger.Info("Launching Wails Dashboard...")
		if err := LaunchApp(id); err != nil {
			logger.Error("Failed to launch Wails app: ", err)
		}

	case 1002:
		logger.Info("Function Id: ", id)
	}
}

func LaunchApp(id int) error {
	if appCmd != nil && appCmd.Process != nil {
		// Already running
		return nil
	}

	if appinfo.Current.BuildType == "development" {
		appCmd = exec.Command("wails", "dev")
		appCmd.Dir = "../app"
		return appCmd.Start()
	}

	appCmd = exec.Command(
		appinfo.Dir.App,
		"-window",
		fmt.Sprintf("%d", id),
	)

	return appCmd.Start()
}
