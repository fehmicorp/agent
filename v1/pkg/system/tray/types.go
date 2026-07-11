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
	Name        string `yaml:"name" json:"name"`
	Title       string `yaml:"title" json:"title"`
	Version     string `yaml:"version" json:"version"`
	Tooltip     string `yaml:"tooltip" json:"tooltip"`
	Icon        string `yaml:"icon" json:"icon"`
	ShowVersion bool   `yaml:"showVersion" json:"showVersion"`
	Menu        []Menu `yaml:"menu" json:"menu"`
}

type Menu struct {
	ID        int    `yaml:"id" json:"id"`
	Title     string `yaml:"title" json:"title"`
	Tooltip   string `yaml:"tooltip" json:"tooltip"`
	FuncId    int    `yaml:"functionId" json:"functionId"`
	Separator bool   `yaml:"separator" json:"separator"`
	Visible   bool   `yaml:"visible" json:"visible"`
	Enabled   bool   `yaml:"enabled" json:"enabled"`
}

type TrayResponse struct {
	Status string `json:"status"`
	Data   struct {
		Tray []Menu `json:"tray"`
	} `json:"data"`
}

func LoadTrayConfig() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	cfg := &Tray{
		Name:        "fehmi-tray-v1",
		Title:       "Fehmi Agent",
		Version:     appinfo.Current.Version,
		Tooltip:     appinfo.Current.Name,
		Icon:        "fav.ico",
		ShowVersion: true,
	}
	resp, err := client.Get(config.URI.API)
	if err != nil {
		logger.Error("Failed to fetch tray config: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logger.Error("server returned status %d", resp.StatusCode)
	}
	var result TrayResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logger.Error("Invalid JSON: %v", err)
	}
	cfg.Menu = result.Data.Tray
	logger.Info("Loaded %d tray menu items", len(cfg.Menu))
}
