package app

import (
	"context"

	"win/internal/config"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ID   int
	Href string
	ctx  context.Context
}

func NewApp(id int) *App {
	return &App{
		ID:   id,
		Href: AppRoute(id),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	config.SetContext(a.ID, ctx)
}

func QuitApp(id int) {

	ctx := config.GetContext(id)

	if ctx == nil {
		return
	}

	config.DeleteContext(id)

	runtime.Quit(ctx)
}

func RedirectApp(id int) {

}

func RunApp(id int) error {

	app := NewApp(id)

	return wails.Run(&options.App{
		Title:       AppTitle(id),
		Width:       AppLayout(id).Width,
		Height:      AppLayout(id).Height,
		StartHidden: AppHidden(id),
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

func AppStartup(app *App) func(context.Context) {
	return func(ctx context.Context) {
		app.startup(ctx)
		if AppOnStartup(app.ID) {
			config.SetContext(app.ID, ctx)
			ShowRoute(ctx, app.Href)
			runtime.LogInfo(
				ctx,
				"Application Started",
			)
		}
	}
}

func AppBeforeClose(app *App) func(context.Context) bool {
	return func(ctx context.Context) bool {
		if AppOnBeforeClose(app.ID) {
			config.DeleteContext(app.ID)
			return false
		}
		return false
	}
}
