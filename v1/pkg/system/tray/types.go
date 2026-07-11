package tray

import "github.com/fehmicorp/agent/v1/cmd/appinfo"

var Current *Tray

type Tray struct {
	Name        string `json:"name,omitempty"`
	Title       string `json:"title,omitempty"`
	Version     string `json:"version,omitempty"`
	Tooltip     string `json:"tooltip,omitempty"`
	Icon        string `json:"icon,omitempty"`
	ShowVersion bool   `json:"showVersion,omitempty"`
	Menu        []Menu `json:"menu,omitempty"`
}

type Menu struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Tooltip   string `json:"tooltip,omitempty"`
	FuncId    int    `json:"functionId,omitempty"`
	Icon      string `json:"icon,omitempty"`
	Separator bool   `json:"separator,omitempty"`
	Visible   bool   `json:"visible,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
}

type TrayResponse struct {
	Status string `json:"status"`
	Data   struct {
		Tray []Menu `json:"tray"`
	} `json:"data"`
}

func DefaultTray() *Tray {
	return &Tray{
		Name:        "fehmi-tray-v1",
		Title:       "Fehmi Agent",
		Version:     appinfo.Current.Version,
		Tooltip:     appinfo.Current.Name,
		Icon:        "fav.ico",
		ShowVersion: true,
		Menu:        []Menu{},
	}
}
