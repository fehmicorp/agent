package tray

import "github.com/fehmicorp/agent/v1/res/logger"

func onExit() {
	logger.Info("Closing Fehmi Agent...")
	if appCmd != nil && appCmd.Process != nil {
		if err := appCmd.Process.Kill(); err != nil {
			logger.Error("Unable to stop app: ", err)
		} else {
			logger.Info("Wails application stopped.")
		}

		_, _ = appCmd.Process.Wait()
		appCmd = nil
	}

	logger.Info("Tray exited.")
}
