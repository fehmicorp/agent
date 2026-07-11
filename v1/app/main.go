package main

import (
	runner "github.com/fehmicorp/agent/v1/app/run"
	"github.com/fehmicorp/agent/v1/cmd/appinfo"
)

// func main() {
// 	go runner.Run(appinfo.Current.Name, 500, 500)
// }

func main() {
	runner.Run(appinfo.Current.Name, 500, 500)
}
