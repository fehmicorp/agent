package internal

import "time"

type Mode string

const (
	Development Mode = "development"
	Beta        Mode = "beta"
	Release     Mode = "stable"
)

// These values are overridden by -ldflags.
var (
	Version   = "0.0.1"
	BuildType = "development"
	BuildDate = ""
	Commit    = "local"
	Branch    = "main"
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

func init() {
	if BuildDate == "" {
		BuildDate = CurrentDate()
	}

	Current = Build{
		Name:      "Fehmi Endpoint Agent",
		Company:   "Fehmi Corporation",
		Version:   Version,
		Build:     BuildDate,
		BuildType: Mode(BuildType),

		Debug:    BuildType == string(Development),
		DevTools: BuildType == string(Development),
		Console:  BuildType == string(Development),
		Admin:    true,
	}
}
