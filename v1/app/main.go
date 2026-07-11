package main

import (
	"github.com/fehmicorp/agent/v1/cmd/appinfo"
)

func main() {
	Run(appinfo.Current.Name, 500, 500)
}
