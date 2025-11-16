package main

import (
	"errors"
	"log"

	"github.com/leonlatsch/go-resolve/internal/config"
	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/service"
)

func main() {
	log.Println("Starting Application")
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	InitializeServiceLocator(conf)

	service, err := createService(conf)
	if err != nil {
		log.Fatalln(err)
	}

	if err := service.Initialize(); err != nil {
		log.Fatalln(err)
	}

	service.ObserveAndUpdateDns()
}

func createService(conf *models.Config) (service.DnsModeService, error) {
	switch conf.Provider {
	case models.ProviderUpdateUrl:
		return &ServiceLocator.UpdateUrlService, nil
	case models.ProviderGoDaddy:
		return &ServiceLocator.GoDaddyService, nil
	case models.ProviderHetzner:
		return &ServiceLocator.HetznerService, nil
	case models.ProviderHetznerCloud:
		return &ServiceLocator.HetznerCloudService, nil
	}

	return nil, errors.New("No service for configured provider")
}
