package registry

import "github.com/fehmicorp/agent/windows/config/types"

func GetTray() *types.Tray {

	if instance == nil {
		return nil
	}

	return &instance.Tray
}
