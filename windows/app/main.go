package main

import (
	"context"

	"github.com/fehmicorp/agent/windows/config/registry"
	"github.com/fehmicorp/agent/windows/config/types"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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

func AppStartup(app *types.NewApp) func(context.Context) {
	return func(ctx context.Context) {
		app.startup(ctx)
		if AppOnStartup(app.ID) {
			registry.SetContext(app.ID, ctx)
			ShowRoute(ctx, app.Href)
			runtime.LogInfo(
				ctx,
				"Application Started",
			)
		}
	}
}

func AppBeforeClose(app *types.NewApp) func(context.Context) bool {
	return func(ctx context.Context) bool {
		if AppOnBeforeClose(app.ID) {
			registry.DeleteContext(app.ID)
			return false
		}
		return false
	}
}
