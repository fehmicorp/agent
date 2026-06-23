package types

type Window struct {
	Id          int    `yaml:"id,omitempty"`
	Title       string `yaml:"title,omitempty"`
	Route       string `yaml:"route,omitempty"`
	Layout      Layout `yaml:"layout,omitempty"`
	Hidden      bool   `yaml:"hide,omitempty"`
	Assets      Assets `yaml:"assets,omitempty"`
	Startup     bool   `yaml:"startup,omitempty"`
	BeforeClose bool   `yaml:"beforeClose,omitempty"`
}

type Layout struct {
	Width  int `yaml:"w,omitempty"`
	Height int `yaml:"h,omitempty"`
}

type Assets struct {
	Icon []byte `yaml:"icon,omitempty"`
}
