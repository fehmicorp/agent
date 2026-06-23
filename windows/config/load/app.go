package load

import (
	"fmt"
	"os"

	"github.com/fehmicorp/agent/windows/config/types"
	"gopkg.in/yaml.v3"
)

func AppConfig() (*types.Window, error) {
	cfgPath, err := os.ReadFile("./yaml/app.yaml")
	var cfg types.Window
	err = yaml.Unmarshal(cfgPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cleanenv failed to load config: %w", err)
	}
	return &cfg, nil
}
