package config

import (
	"errors"
	"os"

	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/serialization"
)

const CONFIG_DIR = "config"
const CONFIG_FILE = CONFIG_DIR + "/" + "godaddy_config.json"
const CONFIG_FILE_MODE = 644

// Loads the config from file and returns it
func LoadConfig() (models.Config, error) {
	var config models.Config

	if !configExists() {
		createEmptyConfig()
		return config, errors.New("No Config. Creating empty config and exiting.")
	}

	buf, err := os.ReadFile(CONFIG_FILE)

	if err != nil {
		return config, err
	}

	if err := serialization.FromJson(string(buf), &config); err != nil {
		return config, err
	}

	return config, nil
}

// Saves SharedConfig to File and applies it to SharedConfig
func SaveConfig(conf models.Config) error {
	confJson, err := serialization.ToJson(conf)
	if err != nil {
		return err
	}

	os.Mkdir(CONFIG_DIR, CONFIG_FILE_MODE) // Ignore err. Just try every time
	if err := os.WriteFile(CONFIG_FILE, []byte(confJson), CONFIG_FILE_MODE); err != nil {
		return err
	}

	return nil
}

func createEmptyConfig() error {
	newEmptyConfig := models.Config{
		ApiKey:    "GODADDY_API_KEY",
		ApiSecret: "GODADDY_API_SECRET",
		Domain:    "YOUR_DOMAIN",
		Hosts:     []string{"HOST1", "HOST2"},
		Interval:  "1h",
	}

	if err := SaveConfig(newEmptyConfig); err != nil {
		return err
	}

	return nil
}

func configExists() bool {
    _, err := os.Stat(CONFIG_FILE)

    return !os.IsNotExist(err)
}
