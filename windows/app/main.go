package main

import (
	"embed"

	"github.com/fehmicorp/agent/windows/debug/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	logger.InitLogger("", "")
	defer logger.Logger.Close()
	if !runas.IsAdmin() {
		logger.Logger.Log("WARN", "Application requires administrative permissions. Requesting UAC elevation...")
		err := runas.RequestElevation()
		if err != nil {
			logger.Logger.Log("ERROR", "UAC prompt escalation rejected or aborted: "+err.Error())
			runtime.Quit(a.ctx)
			return
		}
	}
}
