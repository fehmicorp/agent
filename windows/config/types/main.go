package types

type App struct {
	App    *Application `yaml:"app,omitempty" json:"app,omitempty"`
	Window *Window      `yaml:"window,omitempty" json:"window,omitempty"`
	Tray   *Tray        `yaml:"tray,omitempty" json:"tray,omitempty"`
}
