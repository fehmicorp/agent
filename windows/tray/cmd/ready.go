package main

import (
	assets "github.com/fehmicorp/win_tray/internal/asstes"
	"github.com/fehmicorp/win_tray/internal/config"
	"github.com/getlantern/systray"
)

var cfg *config.Tray

func onReady() {
	if cfg == nil {
		panic("tray config not loaded")
	}
	systray.SetIcon(assets.Favicon)
	systray.SetTitle(cfg.Title)
	systray.SetTooltip(cfg.Tooltip)
}
