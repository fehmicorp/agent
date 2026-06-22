package config

type Database struct {
	Type string `yaml:"type" json:"type"`

	SQLite struct {
		Path string `yaml:"path" json:"path"`
	} `yaml:"sqlite" json:"sqlite"`
}
