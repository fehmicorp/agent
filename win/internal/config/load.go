package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	instance *App
)

func InitialLoad() *App {

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" ||
		func() bool {
			_, err := os.Stat(configPath)
			return os.IsNotExist(err)
		}() {

		exePath, err := os.Executable()
		if err != nil {
			log.Fatal("Could not find executable path")
		}

		configPath = filepath.Join(
			filepath.Dir(exePath),
			"config.yaml",
		)
	}

	var cfg App

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf(
			"can not read config file: %s",
			err.Error(),
		)
	}

	instance = &cfg

	log.Println("Config Path Setting:", configPath)
	log.Printf("Functions Loaded: %d", len(cfg.Functions))
	log.Printf("Tray Loaded: %+v", cfg.Tray)

	return instance
}
