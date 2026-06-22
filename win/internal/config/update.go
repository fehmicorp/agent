package config

type Update struct {
	Enabled       bool   `yaml:"enabled" json:"enabled"`
	CheckInterval string `yaml:"checkInterval" json:"checkInterval"`
	Channel       string `yaml:"channel" json:"channel"`
	AutoInstall   bool   `yaml:"autoInstall" json:"autoInstall"`
}
