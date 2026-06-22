package types

type NewApp struct {
	App  Application `yaml:"app" json:"app"`
	Tray Tray        `yaml:"tray" json:"tray"`
}
