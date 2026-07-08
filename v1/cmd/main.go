package main

import (
	"github.com/fehmicorp/agent/v1/app"
)

func main() {
	// println("Build Type: ", appinfo.Current.BuildType)
	// println("Build Date: ", appinfo.Current.Build)
	// println("Company: ", appinfo.Current.Company)
	// println("Version: ", appinfo.Current.Version)

	go app.Version()
}
