package tray

import (
	"fmt"
	"net/http"
	"os/exec"
	"sync"
	"time"

	"github.com/fehmicorp/agent/v1/cmd/appinfo"
	"github.com/fehmicorp/agent/v1/res/assets"
	"github.com/fehmicorp/agent/v1/res/logger"
	"github.com/getlantern/systray"
)

var (
	appCmd  *exec.Cmd
	appLock sync.Mutex
)

func onReady() {
	if cfg == nil {
		logger.Warn("System tray config not loaded")
		return
	}
	icon, err := assets.Read(cfg.Icon)
	if err == nil {
		systray.SetIcon(icon)
	}
	systray.SetTitle(cfg.Title)
	systray.SetTooltip(fmt.Sprintf("%s\n%s", cfg.Tooltip, cfg.Version))
	populateMenu(cfg)
}

func populateMenu(cfg *Tray) {
	for _, t := range cfg.Menu {
		if t.Separator {
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
			}
		}
		if !t.Enabled {
			item.Disable()
			continue
		}
		go func(menu *systray.MenuItem, id int) {
			for range menu.ClickedCh {
				handleTrayClick(id)
			}
		}(item, t.FuncId)
	}
}

func handleTrayClick(id int) {
	switch id {
	case 9999:
		logger.Warn("Exit menu option clicked.")
		systray.Quit()

	case 1001:
		logger.Info("Launching Dashboard...")
		if err := LaunchApp(id); err != nil {
			logger.Error("Unable to launch app: ", err)
		}

	default:
		logger.Info("Function clicked: ", id)
	}
}

func LaunchApp(id int) error {
	appLock.Lock()
	defer appLock.Unlock()
	if appCmd != nil && appCmd.Process != nil {
		client := &http.Client{
			Timeout: 2 * time.Second,
		}
		req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8051/show", nil)
		if err != nil {
			return err
		}
		resp, err := client.Do(req)
		if err != nil {
			logger.Error("Unable to notify application: ", err)
			return err
		}
		defer resp.Body.Close()
		return nil
	}
	if appinfo.Current.BuildType == "development" {
		appCmd = exec.Command("wails", "dev")
		appCmd.Dir = "../app"
	} else {
		appCmd = exec.Command(
			appinfo.Dir.App,
			"-window",
			fmt.Sprintf("%d", id),
		)
	}
	if err := appCmd.Start(); err != nil {
		appCmd = nil
		return err
	}
	logger.Info("Application started.")
	go func(cmd *exec.Cmd) {
		_ = cmd.Wait()
		appLock.Lock()
		defer appLock.Unlock()
		if appCmd == cmd {
			appCmd = nil
		}
		logger.Info("Application exited.")
	}(appCmd)
	return nil
}
