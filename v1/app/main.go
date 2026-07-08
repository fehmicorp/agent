package main

import (
	"github.com/fehmicorp/agent/v1/cmd/appinfo"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func Version() {
	println("App Name: ", appinfo.Current.Name)
	println("Version: ", appinfo.Current.Version)
}

func main() {
	// Create an instance of the app structure
	app := NewApp()
	Version()

	icon, _ := assets.Read("assets/favicon.ico")

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "app",
		Width:  1024,
		Height: 768,
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
		println("Error:", err.Error())
	}
}
