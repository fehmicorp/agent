package main

import (
	"embed"

	"github.com/fehmicorp/agent/windows/debug/logger"
	"github.com/fehmicorp/agent/windows/utils/runas"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	logger.InitLogger("", "")
	defer logger.Logger.Close()

	// Check if running with administrative rights
	if !runas.IsAdmin() {
		logger.Logger.Log("INFO", "Running as standard user. Attempting elevation request...")
		err := runas.RequestElevation()
		if err != nil {
			logger.Logger.Log("ERROR", "Privilege escalation prompt failed or rejected: "+err.Error())
			// Exit immediately since the user denied permission or elevation failed
			return
		}
	}

	logger.Logger.Log("INFO", "Process verified with Administrative privileges. Starting application context...")

	app := NewApp()
	err := wails.Run(&options.App{
		Title:         "Fehmi Endpoint Agent",
		DisableResize: true,
		Width:         650,
		Height:        500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		logger.Logger.Log("FATAL", "Wails execution encounter: "+err.Error())
		println("Error:", err.Error())
	}
}
