package internal

import "time"

type Mode string

const (
	Development Mode = "development"
	Beta        Mode = "beta"
	Release     Mode = "stable"
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

var (
	Version   = "1.0.0"
	BuildType = "development"
)

func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

var Current = Build{
	Name:      "Fehmi Endpoint Agent",
	Company:   "Fehmi Corporation",
	Version:   Version,
	Build:     CurrentDate(),
	BuildType: Mode(BuildType),

	Debug:    true,
	DevTools: true,
	Console:  true,
	Admin:    true,
}
