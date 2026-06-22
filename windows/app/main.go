package main

import (
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func RunApp(id int) {
	app := NewApp(id)
	return wails.Run(&options.App{
		Title:            AppTitle(id),
		Width:            AppLayout(id).Width,
		Height:           AppLayout(id).Height,
		StartHidden:      AppHidden(id),
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		AssetServer: &assetserver.Options{
			Assets: AppAssets(id),
		},
		OnStartup:     AppStartup(app),
		OnBeforeClose: AppBeforeClose(app),
		Bind: []interface{}{
			app,
		},
	})
}
