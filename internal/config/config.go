package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

type Config struct {
	CurrentUserName string `json:"current_user_name"`
	DbUrl           string `json:"db_url"`
}

const godoDirectory = ".go-do"

const configFileName = "config.json"

func Read() (Config, error) {
	path := getConfigFilePath()

	dat, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("couldn't read config file: %s", err)
	}

	cfg := Config{}
	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshaling config file: %s", err)
	}

	return cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username

	return write(cfg)
}

func (cfg *Config) SetDbUrl(dbUrl string) error {
	cfg.DbUrl = dbUrl

	return write(cfg)
}

func write(cfg *Config) error {
	dat, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error marshaling config data: %s", err)
	}

	path := getConfigFilePath()
	if err := os.WriteFile(path, dat, 0644); err != nil {
		return fmt.Errorf("error writing config file: %s", err)
	}

	return nil
}

func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("no home directory found")
	}

	filePath := path.Join(homeDir, godoDirectory, configFileName)
	return filePath
}
