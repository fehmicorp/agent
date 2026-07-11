package tray

import (
	"net/http"
	"time"

	"github.com/fehmicorp/agent/v1/res/logger"
)

func onExit() {
	logger.Info("Closing Fehmi Agent...")
	appLock.Lock()
	cmd := appCmd
	appLock.Unlock()
	if cmd != nil && cmd.Process != nil {
		logger.Info("Stopping Wails application...")
		client := &http.Client{
			Timeout: 2 * time.Second,
		}
		req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8051/quit", nil)
		if err != nil {
			logger.Error("Unable to create quit request: ", err)
		} else {
			resp, err := client.Do(req)
			if err != nil {
				logger.Error("Unable to send quit request: ", err)
			} else {
				resp.Body.Close()
			}
		}
		done := make(chan error, 1)
		go func() {
			done <- cmd.Wait()
		}()
		select {
		case err := <-done:
			if err != nil {
				logger.Error("Application exited with error: ", err)
			} else {
				logger.Info("Application exited successfully.")
			}
		case <-time.After(5 * time.Second):
			logger.Warn("Application did not exit in time. Terminating...")

			if err := cmd.Process.Kill(); err != nil {
				logger.Error("Unable to terminate application: ", err)
			} else {
				_, _ = cmd.Process.Wait()
				logger.Info("Application terminated.")
			}
		}
		appLock.Lock()
		if appCmd == cmd {
			appCmd = nil
		}
		appLock.Unlock()
	}
	logger.Info("Tray exited.")
}
