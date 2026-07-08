package logs

import (
	"os"

	"github.com/fehmicorp/agent/v1/internal/logger"
)

func Clean() error {

	if logger.Logger != nil {
		logger.Logger.Close()
	}

	err := os.Remove(LogFile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	logger.InitLogger(LogDirectory, "agent.log")

	return nil
}
