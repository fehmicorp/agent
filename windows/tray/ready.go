package main

import (
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
	systray.SetTooltip(cfg.Tooltip)
	go populateMenu(cfg)
}
