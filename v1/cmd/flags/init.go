package internal

type Mode string

const (
	Development Mode = "development"
	Production  Mode = "production"
	Release     Mode = "release"
)

type Build struct {
	Name      string
	Company   string
	Version   string
	Build     string
	BuildType Mode

	Debug     bool
	DevTools  bool
	Console   bool
	Admin     bool
	Profiling bool
}

var Current = Build{
	Name:      "Fehmi Endpoint Agent",
	Company:   "Fehmi Corporation",
	Version:   "1.0.0",
	Build:     "2026.07.08",
	BuildType: Development,

	Debug:    true,
	DevTools: true,
	Console:  true,
	Admin:    true,
}
