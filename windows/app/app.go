package main

import (
	"context"

	"github.com/fehmicorp/agent/windows/debug/logger"
)

type App struct {
	ctx context.Context
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	logger.Logger.Log("INFO", "Running with high integrity context. Spinning up metrics tickers...")
	go a.StartMetric(1000)
}
