package main

import (
	"embed"
	"log"

	"github.com/fehmicorp/agent/windows/debug/logger"
	"github.com/fehmicorp/agent/windows/utils/runas"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

const (
	AppName    = "Fehmi Endpoint Agent"
	AppBuild   = "2026.06.30"
	AppCompany = "Fehmi Corporation"
)

// var (
// 	Version   = appinfo.Version
// 	Build     = appinfo.Build
// 	BuildType = appinfo.BuildType
// )

var signatureSalt = "fehmi-endpoint-windows-agent"

func main() {
	logger.InitLogger("", "")
	defer logger.Logger.Close()
	if err := runas.RequireAdministrator(); err != nil {
		log.Fatal(err)
	}
	app := NewApp()
	err := wails.Run(&options.App{
		Title:         "Fehmi Endpoint Agent",
		DisableResize: true,
		Width:         650,
		Height:        550,
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
