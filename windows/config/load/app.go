package load

import (
	"github.com/fehmicorp/agent/windows/config/types"
	"gopkg.in/yaml.v3"
)

//go:embed yaml/app.yaml
var AppYAML []byte

func AppConfig() ([]types.Window, error) {
	var cfg []types.Window
	err := yaml.Unmarshal(
		AppYAML,
		&cfg,
	)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
