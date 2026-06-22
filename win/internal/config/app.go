package config

type App struct {
	App       Application `yaml:"app" json:"app"`
	Functions []Functions `yaml:"functions" json:"functions"`
	Tray      Tray        `yaml:"tray" json:"tray"`
}
