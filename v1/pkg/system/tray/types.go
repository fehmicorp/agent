package tray

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fehmicorp/agent/v1/cmd/appinfo"
	"github.com/fehmicorp/agent/v1/cmd/config"
	"github.com/fehmicorp/agent/v1/res/logger"
)

var cfg *Tray

type Tray struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Version     string `json:"version"`
	Tooltip     string `json:"tooltip"`
	Icon        string `json:"icon"`
	ShowVersion bool   `json:"showVersion"`
	Menu        []Menu `json:"menu"`
}

type Menu struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Tooltip   string `json:"tooltip,omitempty"`
	FuncId    int    `json:"functionId,omitempty"`
	Icon      string `json:"icon,omitempty"`
	Separator bool   `json:"separator,omitempty"`
	Visible   bool   `json:"visible,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
}

func LoadTrayConfig() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	cfg = &Tray{
		Name:        "fehmi-tray-v1",
		Title:       "Fehmi Agent",
		Version:     appinfo.Current.Version,
		Tooltip:     appinfo.Current.Name,
		Icon:        "fav.ico",
		ShowVersion: true,
	}
	resp, err := client.Get(config.URI.API)
	if err != nil {
		logger.Error("Failed to fetch tray config: ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logger.Error("server returned status: ", resp.StatusCode)
	}
	var result TrayResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logger.Error("Invalid JSON:", err)
	}
	cfg.Menu = result.Data.Tray
	logger.Info("Loaded tray menu items: ", len(cfg.Menu))
}
