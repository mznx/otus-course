package config

import (
	"encoding/json"
	"os"
)

func ReadConfig() *Config {
	c := &Config{}

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "config.json"
	}

	if err := readJson(configPath, c); err != nil {
		panic(err)
	}

	return c
}

func readJson(filepath string, conf any) error {
	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(fileBytes, conf); err != nil {
		return err
	}

	return nil
}
