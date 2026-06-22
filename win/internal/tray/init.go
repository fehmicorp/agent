package tray

import (
	"errors"
	"win/internal/config"
)

func Init() error {

	tryCfg = config.GetTray()

	if tryCfg == nil {
		return errors.New("tray config not found")
	}

	return nil
}
