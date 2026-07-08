package main

import (
	"githum.com/fehmicorp/v1/app"
	"githum.com/fehmicorp/v1/cmd/internal"
)

func main() {
	println("Build Type: ", internal.Current.BuildType)
	println("Build Date: ", internal.Current.Build)
	println("Company: ", internal.Current.Company)
	println("Version: ", internal.Current.Version)

	go app.Run()
}
