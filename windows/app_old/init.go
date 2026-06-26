package main

import (
	"github.com/fehmicorp/agent/windows/config/load"
	"github.com/fehmicorp/agent/windows/config/types"
)

var Windows []types.Window

func Init() {
	var err error
	Windows, err = load.AppConfig()
	if err != nil {
		panic(err)
	}
	RegisterMany(Windows)
}
