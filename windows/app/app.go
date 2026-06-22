package main

import (
	"context"

	"github.com/fehmicorp/agent/windows/config/registry"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		Id:   id,
		Href: AppRoute(id),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	registry.SetContext(a.Id, ctx)
}
