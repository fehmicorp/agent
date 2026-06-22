package config

type Logging struct {
	Level      string `yaml:"level" json:"level"`
	File       string `yaml:"file" json:"file"`
	MaxSizeMB  int    `yaml:"maxSizeMB" json:"maxSizeMB"`
	MaxBackups int    `yaml:"maxBackups" json:"maxBackups"`
	MaxAgeDays int    `yaml:"maxAgeDays" json:"maxAgeDays"`
}
