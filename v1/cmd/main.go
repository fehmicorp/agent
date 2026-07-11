package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/fehmicorp/agent/v1/cmd/appinfo"
	v "github.com/fehmicorp/agent/v1/cmd/version"
	"github.com/fehmicorp/agent/v1/pkg/system/notification"
	"github.com/fehmicorp/agent/v1/pkg/system/tray"
	"github.com/fehmicorp/agent/v1/res/logger"
)

func main() {
	logger.Init("", "")
	dev := flag.Bool("dev", false, "Run Wails development")
	build := flag.Bool("build", false, "Build Wails application")
	clean := flag.Bool("clean", true, "Clean build output")
	flag.Parse()
	switch {
	case *dev:
		v.BuildType = "development"
	case *build:
		v.BuildType = "beta"
	case *clean:
		v.BuildType = "stable"
	}
	switch v.BuildType {
	case "development", "beta", "stable":
	default:
		logger.Warn("invalid build type:", v.BuildType)
	}
	appinfo.Reload()
	notification.Register("Fehmi Endpoint Agent")
	logger.Info("----------------------------------------")
	logger.Info("Application :", appinfo.Current.Name)
	logger.Info("Version     :", appinfo.Current.Version)
	logger.Info("Build Type  :", appinfo.Current.BuildType)
	logger.Info("Build Date  :", appinfo.Current.Build)
	logger.Info("----------------------------------------")
	tray.RunTray()
	// switch {
	// case *dev:
	// 	if err := WailsDev(); err != nil {
	// 		logger.Fatal(err)
	// 	}
	// case *build:
	// 	if err := WailsBuild(*clean); err != nil {
	// 		logger.Fatal(err)
	// 	}
	// case *clean:
	// 	if err := WailsBuild(*clean); err != nil {
	// 		logger.Fatal(err)
	// 	}
	// }
}

func WailsDev() error {

	args := []string{
		"dev",
	}

	cmd := exec.Command("wails", args...)
	cmd.Dir = "../app"

	cmd.Env = append(os.Environ(),
		"APP_VERSION="+appinfo.Current.Version,
		"APP_BUILD_TYPE="+string(appinfo.Current.BuildType),
		"APP_BUILD_DATE="+appinfo.Current.Build,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func WailsBuild(clean bool) error {

	args := []string{
		"build",
	}

	if clean {
		args = append(args, "-clean")
	}

	args = append(args,
		"-ldflags",
		fmt.Sprintf(
			"-X github.com/fehmicorp/agent/v1/cmd/version.Version=%s -X github.com/fehmicorp/agent/v1/cmd/version.BuildType=%s -X github.com/fehmicorp/agent/v1/cmd/version.BuildDate=%s",
			appinfo.Current.Version,
			appinfo.Current.BuildType,
			appinfo.Current.Build,
		),
	)

	cmd := exec.Command("wails", args...)
	cmd.Dir = "../app"

	cmd.Env = append(os.Environ(),
		"APP_VERSION="+appinfo.Current.Version,
		"APP_BUILD_TYPE="+string(appinfo.Current.BuildType),
		"APP_BUILD_DATE="+appinfo.Current.Build,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
