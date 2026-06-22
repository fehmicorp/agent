package main

import (
	"fmt"

	assets "github.com/fehmicorp/agent/windows/config/assets"
	"github.com/fehmicorp/agent/windows/config/types"
	"github.com/getlantern/systray"
)

var cfg *types.Tray

func onReady() {
	if cfg == nil {
		panic("tray config not loaded")
	}
	systray.SetIcon(assets.Favicon)
	systray.SetTitle(cfg.Title)
	systray.SetTooltip(fmt.Sprintf("%s\n%s", cfg.Tooltip, cfg.Version))
	go populateMenu(cfg)
}
