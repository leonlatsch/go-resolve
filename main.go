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

	service, err := createDnsService(conf)
	if err != nil {
		log.Fatalln(err)
	}

	if err := service.Initialize(); err != nil {
		log.Fatalln(err)
	}

	ipObserverService := &ServiceLocator.IpObserverService

	log.Println("Running for " + conf.Provider)

	ipObserverService.ObserveIp(func(ip string) {
		log.Println("New IP: " + ip + " | Notifying " + conf.Provider)
		err := service.UpdateDns(ip)
		if err != nil {
			log.Println("Not caching ip: ", err)
		} else {
			log.Println("Successfully updated all records. Caching " + ip)
			ipObserverService.LastIp = ip
		}
	})
}

func createDnsService(conf *models.Config) (service.DnsModeService, error) {
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
