package main

import (
	"github.com/fehmicorp/agent/v1/app"
)

func main() {
	// println("Build Type: ", internal.Current.BuildType)
	// println("Build Date: ", internal.Current.Build)
	// println("Company: ", internal.Current.Company)
	// println("Version: ", internal.Current.Version)

	go app.Run()
}
