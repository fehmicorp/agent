package logs

import (
	"path/filepath"

	"github.com/fehmicorp/agent/v1/internal/logger"
)

var (
	LogDirectory string
	LogFile      string
)

func Init(path string) {

	LogDirectory = path

	if LogDirectory == "" {
		LogDirectory = "."
	}

	LogFile = filepath.Join(LogDirectory, "agent.log")

	logger.InitLogger(LogDirectory, "agent.log")
}
