package types

type Window struct {
	Id          int
	Title       string
	Route       string
	Layout      Layout
	Hidden      bool
	Assets      Assets
	Startup     bool
	BeforeClose bool
}

type Layout struct {
	Width  int
	Height int
}

type Assets struct {
	Icon []byte
}
