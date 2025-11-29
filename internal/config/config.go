package config

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/serialization"
)

const (
	CONFIG_DIR       = "config"
	CONFIG_FILE      = CONFIG_DIR + "/" + "config.json"
	CONFIG_FILE_MODE = 0o644
)

// Get the config as pointer. Value updates when file changes
func GetConfig() (*models.Config, error) {
	if !configExists() {
		createEmptyConfig()
		return nil, errors.New("no Config. Creating empty config and exiting")
	}

	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	observeConfigFile(func(c models.Config) {
		config = c
	})

	return &config, nil
}

// Saves SharedConfig to File and applies it to SharedConfig
func SaveConfig(conf models.Config) error {
	confJson, err := serialization.ToJson(conf)
	if err != nil {
		return err
	}

	os.Mkdir(CONFIG_DIR, 0o755) // Ignore err. Just try every time
	if err := os.WriteFile(CONFIG_FILE, []byte(confJson), CONFIG_FILE_MODE); err != nil {
		return err
	}

	return nil
}

// Loads the config from file and returns it
func loadConfig() (models.Config, error) {
	var config models.Config
	buf, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		return config, err
	}

	if err := serialization.FromJson(string(buf), &config); err != nil {
		return config, err
	}

	return config, nil
}

func createEmptyConfig() error {
	if err := SaveConfig(models.EmptyConfig); err != nil {
		return err
	}

	return nil
}

func configExists() bool {
	_, err := os.Stat(CONFIG_FILE)

	return !os.IsNotExist(err)
}

func observeConfigFile(onChange func(models.Config)) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
		return
	}

	if err := watcher.Add(CONFIG_DIR); err != nil {
		log.Println(err)
		return
	}

	go func() {
		defer watcher.Close()

		for {
			e := <-watcher.Events
			if !strings.HasSuffix(e.Name, CONFIG_FILE) {
				continue
			}

			log.Println(e.Name + " has changed. Reloading")

			newConf, err := loadConfig()
			if err != nil {
				log.Println(err)
				continue
			}

			onChange(newConf)
		}
	}()
}
