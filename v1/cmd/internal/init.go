package internal

import (
	"time"

	v "github.com/fehmicorp/agent/v1/cmd/internal/version"
)

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

func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

var Current Build

func init() {

	if v.BuildDate == "" {
		v.BuildDate = CurrentDate()
	}

	Current = Build{
		Name:      "Fehmi Endpoint Agent",
		Company:   "Fehmi Corporation",
		Version:   v.Version,
		Build:     v.BuildDate,
		BuildType: Mode(v.BuildType),

		Debug:     false,
		DevTools:  false,
		Console:   false,
		Admin:     true,
		Profiling: false,
	}

	switch Current.BuildType {

	case Development:
		Current.Debug = true
		Current.DevTools = true
		Current.Console = true
		Current.Admin = false
		Current.Profiling = true

	case Beta:
		Current.Debug = false
		Current.DevTools = false
		Current.Console = true
		Current.Admin = true
		Current.Profiling = false

	case Release:
		Current.Debug = false
		Current.DevTools = false
		Current.Console = false
		Current.Admin = true
		Current.Profiling = false
	}
}
