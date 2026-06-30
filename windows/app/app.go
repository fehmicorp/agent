package main

import (
	"context"

	"github.com/fehmicorp/agent/windows/debug/logger"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	logger.Logger.Log("INFO", "Running with high integrity context. Spinning up metrics tickers...")
	go a.StartMetric(1000)
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
