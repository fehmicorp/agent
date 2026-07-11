package config

import "github.com/fehmicorp/agent/v1/cmd/appinfo"

type Uris struct {
	API        string
	CDN        string
	Web        string
	Utm_Source string
}

var URI = &Uris{
	API:        getAPI(),
	CDN:        "https://cdn.fehmicorp.in/v1/agnt",
	Web:        "https://fehmicorp.in/",
	Utm_Source: "v1_agent",
}

func getAPI() string {
	if appinfo.Current.BuildType == "development" {
		return "http://localhost:8050"
	}

	return "https://agent.fehmicorp.in/v1/api/"
}
