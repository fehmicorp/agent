package runner

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func Run() error {

	a := NewApp()

	return wails.Run(&options.App{

		Title: "Fehmi Endpoint Agent",

		Width:  1024,
		Height: 768,

		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		OnStartup: a.startup,

		Bind: []interface{}{
			a,
		},
	})
}
