package config

type Tray struct {
	Name        string `yaml:"name" json:"name"`
	Title       string `yaml:"title" json:"title"`
	Version     string `yaml:"version" json:"version"`
	Tooltip     string `yaml:"tooltip" json:"tooltip"`
	Icon        string `yaml:"icon" json:"icon"`
	ShowVersion bool   `yaml:"showVersion" json:"showVersion"`

	Menu []Menu `yaml:"menu" json:"menu"`
}

type Menu struct {
	ID        int    `yaml:"id" json:"id"`
	Title     string `yaml:"title" json:"title"`
	Tooltip   string `yaml:"tooltip" json:"tooltip"`
	FuncId    int    `yaml:"functionId" json:"functionId"`
	Separator bool   `yaml:"separator" json:"separator"`
	Visible   bool   `yaml:"visible" json:"visible"`
	Enabled   bool   `yaml:"enabled" json:"enabled"`
}
