package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	runner "github.com/fehmicorp/agent/v1/app/run"
	"github.com/fehmicorp/agent/v1/cmd/appinfo"
	v "github.com/fehmicorp/agent/v1/cmd/version"
)

func main() {

	version := flag.String("version", "", "Application version")
	buildType := flag.String("buildtype", "", "development|beta|stable")
	build := flag.Bool("build", false, "Build Wails application")

	flag.Parse()

	if *version != "" {
		v.Version = *version
	}

	if *buildType != "" {
		v.BuildType = *buildType
	}

	// Refresh appinfo after updating version variables.
	appinfo.Reload()

	if *build {
		if err := buildWails(); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := runner.Run(
		appinfo.Current.Name,
		500,
		500,
	); err != nil {
		log.Fatal(err)
	}
}

func buildWails() error {

	args := []string{
		"build",
		"-clean",
		"-ldflags",
		"-X github.com/fehmicorp/agent/v1/cmd/version.Version=" + appinfo.Current.Version +
			" -X github.com/fehmicorp/agent/v1/cmd/version.BuildType=" + string(appinfo.Current.BuildType),
	}

	cmd := exec.Command("wails", args...)
	cmd.Dir = "../app"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
