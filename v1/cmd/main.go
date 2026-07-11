package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fehmicorp/agent/v1/cmd/appinfo"
	v "github.com/fehmicorp/agent/v1/cmd/version"
)

func main() {

	version := flag.String("v", v.Version, "Version")
	buildType := flag.String("buildtype", v.BuildType, "development|beta|stable")

	dev := flag.Bool("dev", false, "Run Wails development")
	build := flag.Bool("build", false, "Build Wails application")
	clean := flag.Bool("clean", true, "Clean build output")

	flag.Parse()

	// ------------------------------------------------------------------
	// Apply CLI values
	// ------------------------------------------------------------------

	v.Version = strings.TrimSpace(*version)
	println("Version: ", strings.TrimSpace(*version))
	v.BuildType = strings.ToLower(strings.TrimSpace(*buildType))

	switch v.BuildType {
	case "development", "beta", "stable":
	default:
		log.Fatalf("invalid build type: %s", v.BuildType)
	}

	appinfo.Reload()

	fmt.Println("----------------------------------------")
	fmt.Println("Application :", appinfo.Current.Name)
	fmt.Println("Version     :", appinfo.Current.Version)
	fmt.Println("Build Type  :", appinfo.Current.BuildType)
	fmt.Println("Build Date  :", appinfo.Current.Build)
	fmt.Println("----------------------------------------")

	switch {

	case *dev:
		if err := wailsDev(); err != nil {
			log.Fatal(err)
		}

	case *build:
		if err := wailsBuild(*clean); err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Println()
		fmt.Println("Nothing to do.")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  go run . -dev")
		fmt.Println("  go run . -build")
		fmt.Println("  go run . -build -version=1.2.1 -buildtype=stable")
	}
}

func wailsDev() error {

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

func wailsBuild(clean bool) error {

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
