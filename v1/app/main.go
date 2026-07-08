package main

import (
	runner "github.com/fehmicorp/agent/v1/app/run"
	"github.com/fehmicorp/agent/v1/cmd/appinfo"
)

//	func Version() {
//		println("App Name: ", appinfo.Current.Name)
//		println("Version: ", appinfo.Current.Version)
//	}
func main() {
	go runner.Run(appinfo.Current.Name, 500, 500)
}
