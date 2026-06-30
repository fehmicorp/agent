package main

import (
	"embed"

	"github.com/fehmicorp/agent/windows/debug/logger"
	"github.com/fehmicorp/agent/windows/utils/runas"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func (a *App) main() {
	logger.InitLogger("", "")
	defer logger.Logger.Close()
	if !runas.IsAdmin() {
		logger.Logger.Log("WARN", "Application requires administrative permissions. Requesting UAC elevation...")
		err := runas.RequestElevation()
		if err != nil {
			logger.Logger.Log("ERROR", "UAC prompt escalation rejected or aborted: "+err.Error())
			runtime.Quit(a.ctx)
			return
		}
	}
	err := wails.Run(&options.App{
		Title:         "Fehmi Endpoint Agent",
		DisableResize: true,
		Width:         650,
		Height:        500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        a.startup,
		Bind: []interface{}{
			a,
		},
	})
	if err != nil {
		logger.Logger.Log("FATAL", "Wails execution encounter: "+err.Error())
		println("Error:", err.Error())
	}
}
