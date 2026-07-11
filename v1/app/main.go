package main

import (
	"github.com/fehmicorp/agent/v1/cmd/appinfo"
	"github.com/fehmicorp/agent/v1/res/assets"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func main() {
	Run(appinfo.Current.Name, 500, 500)
}

func Run(title string, w int, h int) error {
	a := NewApp()
	return wails.Run(&options.App{
		Title:  title,
		Width:  w,
		Height: h,
		AssetServer: &assetserver.Options{
			Assets: assets.FS,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        a.startup,
		Bind: []interface{}{
			a,
		},
	})
}
