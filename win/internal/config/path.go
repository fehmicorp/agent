package config

type Paths struct {
	Data      string `yaml:"data" json:"data"`
	Logs      string `yaml:"logs" json:"logs"`
	Cache     string `yaml:"cache" json:"cache"`
	Temp      string `yaml:"temp" json:"temp"`
	Downloads string `yaml:"downloads" json:"downloads"`
}
