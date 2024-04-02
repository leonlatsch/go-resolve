package config

import (
	"log"
	"os"

	"github.com/leonlatsch/go-resolve/models"
	"github.com/leonlatsch/go-resolve/serialization"
)
const CONFIG_DIR = "config"
const CONFIG_FILE = CONFIG_DIR + "/" + "godaddy_config.json" 
const CONFIG_FILE_MODE = 0644

// In memory config. DON'T modify directly. Use config.SaveConfig()
var SharedConfig models.Config

// Loads the config in SharedConfig
func LoadConfig() error {
	if !configExists() {
		createEmptyConfig()
		log.Fatalln("No Config. Creating empty config and exiting.")
	}

	buf, err := os.ReadFile(CONFIG_FILE)

	if err != nil {
		return err
	}

	if err := serialization.FromJson(string(buf), &SharedConfig); err != nil {
		return err
	}

	return nil
}

// Saves SharedConfig to File and applies it to SharedConfig
func SaveConfig(conf models.Config) error {
	confJson, err := serialization.ToJson(conf)
	if err != nil {
		return err
	}
    
    os.Mkdir(CONFIG_DIR, CONFIG_FILE_MODE)// Ignore err. Just try every time
	if err := os.WriteFile(CONFIG_FILE, []byte(confJson), CONFIG_FILE_MODE); err != nil {
		return err
	}

	SharedConfig = conf

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
	if _, err := os.Stat(CONFIG_FILE); os.IsNotExist(err) {
		return false
	}

	return true
}
