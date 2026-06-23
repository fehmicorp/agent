package load

import (
	"github.com/fehmicorp/agent/windows/config"
	"github.com/fehmicorp/agent/windows/config/types"
	"gopkg.in/yaml.v3"
)

func AppConfig() ([]types.Window, error) {
	var cfg []types.Window
	err := yaml.Unmarshal(
		config.AppYAML,
		&cfg,
	)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
