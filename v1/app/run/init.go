package runner

import (
	"github.com/fehmicorp/agent/v1/res/assets"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func Run() error {
	a := NewApp()
	icon, _ := assets.Read("assets/favicon.ico")
	return wails.Run(&options.App{

		Title: "Fehmi Endpoint Agent",

		Width:  1024,
		Height: 768,

		AssetServer: &assetserver.Options{
			Assets: icon,
		},

		OnStartup: a.startup,

		Bind: []interface{}{
			a,
		},
	})
}
