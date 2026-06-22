package config

type Server struct {
	Endpoint string `yaml:"endpoint" json:"endpoint"`
	APIKey   string `yaml:"apiKey" json:"apiKey"`

	Websocket struct {
		Enabled bool   `yaml:"enabled" json:"enabled"`
		URL     string `yaml:"url" json:"url"`
	} `yaml:"websocket" json:"websocket"`
}
