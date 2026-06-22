package main

import (
	"context"

	"github.com/fehmicorp/agent/windows/config/registry"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp(id int) *App {
	return &App{
		ID:   id,
		Href: AppRoute(id),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	registry.SetContext(a.ID, ctx)
}

func QuitApp(id int) {
	ctx := registry.GetContext(id)
	if ctx == nil {
		return
	}
	registry.DeleteContext(id)
	runtime.Quit(ctx)
}
