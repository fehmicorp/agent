package main

import (
	"context"
	"fmt"

	"github.com/fehmicorp/agent/v1/cmd/flags"
	"github.com/fehmicorp/agent/v1/internal/assets"
	"github.com/fehmicorp/agent/v1/internal/logger"
	"github.com/fehmicorp/agent/v1/internal/runas"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

const (
	AppName    = "Fehmi Endpoint Agent"
	AppCompany = "Fehmi Corporation"
	AppBuild   = "2026.06.30"
)

var salt = "fehmi-endpoint-windows-agent"

func Run() error {

	//--------------------------------------------------
	// Build Flags
	//--------------------------------------------------

	flags.ProductionConfig()

	//--------------------------------------------------
	// Logger
	//--------------------------------------------------

	if err := logger.Init("", "agent.log"); err != nil {
		return err
	}

	defer logger.Close()

	logger.Info("========================================")
	logger.Info(AppName)
	logger.Info("Build      : ", AppBuild)
	logger.Info("Mode       : ", flags.Current.BuildType)
	logger.Info("========================================")

	//--------------------------------------------------
	// Assets
	//--------------------------------------------------

	icon, err := assets.Read("fav.ico")
	if err != nil {
		return fmt.Errorf("icon: %w", err)
	}

	//--------------------------------------------------
	// Administrator
	//--------------------------------------------------

	if flags.Current.Admin {

		if err := runas.RequireAdministrator(); err != nil {
			return err
		}
	}

	//--------------------------------------------------
	// Application
	//--------------------------------------------------

	app := NewApp()

	//--------------------------------------------------
	// Wails
	//--------------------------------------------------

	err = wails.Run(&options.App{

		Title: AppName,

		Width:  650,
		Height: 550,

		MinWidth:  650,
		MinHeight: 550,

		DisableResize: true,

		BackgroundColour: &options.RGBA{
			R: 27,
			G: 38,
			B: 54,
			A: 1,
		},

		AssetServer: &assetserver.Options{
			Assets: assets.FS,
		},

		Windows: &options.Windows{
			Icon: icon,
		},

		OnStartup: app.startup,

		OnShutdown: func(ctx context.Context) {
			logger.Info("Application Shutdown")
		},

		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Application Closed")

	return nil
}
