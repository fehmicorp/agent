package load

import (
	"github.com/fehmicorp/agent/windows/config/assets"
	"github.com/fehmicorp/agent/windows/config/types"
)

var Windows = []types.Window{
	{
		Id:    1001,
		Title: "Dashboard",
		Route: "dashboard",
		Layout: types.Layout{
			Width:  1280,
			Height: 800,
		},
		Hidden: false,
		Assets: types.Assets{
			Icon: assets.Favicon,
		},
		Startup:     true,
		BeforeClose: true,
	},
	{
		Id:    1002,
		Title: "Scan",
		Route: "scan",
		Layout: types.Layout{
			Width:  500,
			Height: 500,
		},
		Hidden: false,
		Assets: types.Assets{
			Icon: assets.Favicon,
		},
		Startup:     true,
		BeforeClose: true,
	},
	{
		Id:    1003,
		Title: "Agent",
		Route: "agent",
		Layout: types.Layout{
			Width:  500,
			Height: 500,
		},
		Hidden: false,
		Assets: types.Assets{
			Icon: assets.Favicon,
		},
		Startup:     true,
		BeforeClose: true,
	},
	{
		Id:    1004,
		Title: "Logs",
		Route: "logs",
		Layout: types.Layout{
			Width:  500,
			Height: 500,
		},
		Hidden: false,
		Assets: types.Assets{
			Icon: assets.Favicon,
		},
		Startup:     true,
		BeforeClose: true,
	},
	{
		Id:    1005,
		Title: "Software Catalogues",
		Route: "software",
		Layout: types.Layout{
			Width:  500,
			Height: 500,
		},
		Hidden: false,
		Assets: types.Assets{
			Icon: assets.Favicon,
		},
		Startup:     true,
		BeforeClose: true,
	},
	{
		Id:    1006,
		Title: "Request Support",
		Route: "support",
		Layout: types.Layout{
			Width:  500,
			Height: 500,
		},
		Hidden: false,
		Assets: types.Assets{
			Icon: assets.Favicon,
		},
		Startup:     true,
		BeforeClose: true,
	},
	{
		Id:    1007,
		Title: "Admin",
		Route: "admin",
		Layout: types.Layout{
			Width:  500,
			Height: 500,
		},
		Hidden: false,
		Assets: types.Assets{
			Icon: assets.Favicon,
		},
		Startup:     true,
		BeforeClose: true,
	},
}
