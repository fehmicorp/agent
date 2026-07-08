package app

import "github.com/fehmicorp/agent/v1/cmd/internal"

func Run() {
	println("Build Type: ", internal.Current.BuildType)
	println("Build Date: ", internal.Current.Build)
	println("Company: ", internal.Current.Company)
	println("Version: ", internal.Current.Version)
}
