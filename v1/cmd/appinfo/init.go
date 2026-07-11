package appinfo

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	v "github.com/fehmicorp/agent/v1/cmd/version"
)

type Mode string

type DirStruct struct {
	Base      string
	Installer string
	App       string
	Logs      string
	Config    string
	Data      string
}

var Dir = &DirStruct{}

func InitDirectories() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	Dir.Base = filepath.Dir(exe)

	switch runtime.GOOS {
	case "windows":
		Dir.Installer = Dir.Base
		Dir.App = filepath.Join(Dir.Base, "app.exe")

	case "darwin":
		// Inside MyApp.app/Contents/MacOS
		Dir.Installer = filepath.Dir(filepath.Dir(Dir.Base))
		Dir.App = filepath.Join(Dir.Base, "app")

	case "linux":
		Dir.Installer = Dir.Base
		Dir.App = filepath.Join(Dir.Base, "app")
	}

	Dir.Logs = filepath.Join(Dir.Base, "logs")
	Dir.Config = filepath.Join(Dir.Base, "config")
	Dir.Data = filepath.Join(Dir.Base, "data")

	return nil
}

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
