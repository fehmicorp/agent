package load

import (
	"fmt"

	"github.com/fehmicorp/agent/windows/config/types"
	"github.com/ilyakaznacheev/cleanenv"
)

func AppConfig() (*types.Window, error) {
	cfgPath := "./yaml/app.yaml"
	var cfg types.Window
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cleanenv failed to load config: %w", err)
	}
	return &cfg, nil
}
