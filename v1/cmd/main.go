package main

import internal "githum.com/fehmicorp/v1/cmd/flags"

var (
	Version     = internal.Current.Version
	DisplayName = "Fehmi Endpoint Agent"
	PackageName = "in.fehemicorp.agent"
	Author      = "Fehmi Corporation"
	Website     = "https://www.fehmicorp.in"
)

func main() {
	println(DisplayName)
	println("Version: ", Version)
}
