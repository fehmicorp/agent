package tray

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fehmicorp/agent/v1/cmd/config"
	"github.com/fehmicorp/agent/v1/pkg/db/sqlite"
	"github.com/fehmicorp/agent/v1/res/logger"
)

var store *sqlite.SQLiteStore

func InitialSQLite(dir string) error {
	store = sqlite.NewSQLiteStore("agent.db", "tray")
	_, err := store.Init(dir, []sqlite.TblQuery{
		{
			Key:        "key",
			Type:       "TEXT",
			Preference: "PRIMARY KEY",
		},
		{
			Key:  "value",
			Type: "TEXT",
		},
	})
	return err
}

func SaveTray(cfg *Tray) error {
	if store == nil {
		return fmt.Errorf("tray sqlite not initialized")
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return store.Set("tray", string(data))
}

func LoadTray() (*Tray, error) {
	if store == nil {
		return nil, fmt.Errorf("tray sqlite not initialized")
	}
	value, err := store.Get("tray")
	if err != nil {
		return nil, err
	}
	if value == "" {
		return DefaultTray(), nil
	}
	cfg := DefaultTray()
	if err := json.Unmarshal([]byte(value), cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func DownloadTray() (*Tray, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(config.URI.API)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned %d", resp.StatusCode)
	}
	var result TrayResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	cfg := DefaultTray()
	cfg.Menu = result.Data.Tray
	return cfg, nil
}

func RefreshTray() (*Tray, error) {
	cfg, err := DownloadTray()
	if err != nil {
		logger.Warn("Unable to fetch tray from server. Loading local cache.")
		return LoadTray()
	}
	if err := SaveTray(cfg); err != nil {
		logger.Warn("Unable to cache tray:", err)
	}
	return cfg, nil
}

func LoadTrayConfig() error {
	cfg, err := RefreshTray()
	if err != nil {
		return err
	}
	Current = cfg
	logger.Info("Loaded tray menu items:", len(Current.Menu))
	return nil
}
