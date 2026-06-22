package types

type Window struct {
	Id          int    `yaml:"id"`
	Title       string `yaml:"title"`
	Route       string `yaml:"route"`
	Layout      Layout `yaml:"layout"`
	Hidden      bool   `yaml:"hide"`
	Assets      Assets `yaml:"assets"`
	Startup     bool   `yaml:"startup"`
	BeforeClose bool   `yaml:"beforeClose"`
}

type Layout struct {
	Width  int `yaml:"w"`
	Height int `yaml:"h"`
}

type Assets struct {
	icon string `yaml:"icon"`
}
