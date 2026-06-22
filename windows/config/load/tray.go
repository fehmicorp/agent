package load

import (
	"fmt"

	"github.com/fehmicorp/agent/windows/config/types"
	"github.com/ilyakaznacheev/cleanenv"
)

func TrayConf() (*types.Tray, error) {
	cfgPath := "./yaml/tray.yaml"
	var cfg types.Tray
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cleanenv failed to load config: %w", err)
	}
	return &cfg, nil
}
