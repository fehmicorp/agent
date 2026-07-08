package main

import (
	runner "github.com/fehmicorp/agent/v1/app/run"
	"github.com/fehmicorp/agent/v1/cmd/appinfo"
)

func main() {
	// println("Build Type: ", appinfo.Current.BuildType)
	// println("Build Date: ", appinfo.Current.Build)
	// println("Company: ", appinfo.Current.Company)
	// println("Version: ", appinfo.Current.Version)

	go runner.Run(appinfo.Current.Name, 500, 500)
}
