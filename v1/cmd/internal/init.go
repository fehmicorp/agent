package internal

import (
	"time"

	"github.com/fehmicorp/agent/v1/cmd/internal/version"
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

	Commit string
	Branch string
}

func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

var Current Build

func init() {

	buildDate := version.BuildDate
	if buildDate == "" {
		buildDate = CurrentDate()
	}

	Current = Build{
		Name:      "Fehmi Endpoint Agent",
		Company:   "Fehmi Corporation",
		Version:   version.Version,
		Build:     buildDate,
		BuildType: Mode(version.BuildType),

		Debug:     version.BuildType == string(Development),
		DevTools:  version.BuildType == string(Development),
		Console:   version.BuildType == string(Development),
		Admin:     true,
		Profiling: version.BuildType == string(Development),

		Commit: version.Commit,
		Branch: version.Branch,
	}
}
