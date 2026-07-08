package appinfo

import (
	"time"

	v "github.com/fehmicorp/agent/v1/cmd/version"
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

func Reload() {

	buildDate := v.BuildDate
	if buildDate == "" {
		buildDate = CurrentDate()
	}

	Current = Build{
		Name:      "Fehmi Endpoint Agent",
		Company:   "Fehmi Corporation",
		Version:   v.Version,
		Build:     buildDate,
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
		Current.Console = true
		Current.Admin = true

	case Release:
		Current.Admin = true

	default:
		// Fallback to development
		Current.BuildType = Development
		Current.Debug = true
		Current.DevTools = true
		Current.Console = true
		Current.Admin = false
		Current.Profiling = true
	}
}
func init() {
	Reload()
}
