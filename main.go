package main

import (
	"errors"
	"log"

	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/config"
	"github.com/leonlatsch/go-resolve/internal/godaddy"
	"github.com/leonlatsch/go-resolve/internal/hetzner"
	"github.com/leonlatsch/go-resolve/internal/http"
	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/service"
)

func main() {
	log.Println("Starting Application")
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	httpClient := http.RealHttpClient{}

	service, err := createService(conf, &httpClient)
	if err != nil {
		log.Fatalln(err)
	}

	service.ObserveAndUpdateDns()
}

func createService(conf *models.Config, httpClient http.HttpClient) (service.DnsModeService, error) {
	if conf.Provider == models.ProviderUpdateUrl {
		updateUrlService := service.UpdateUrlService{
			Config:       conf,
			UpdateUrlApi: &api.UpdateUrlApiImpl{Config: conf, HttpClient: httpClient},
			IpObserver: service.IpObserverService{
				IpApi:  &api.IpApiImpl{HttpClient: httpClient},
				Config: conf,
			},
		}

		return &updateUrlService, nil

	}

	if conf.Provider == models.ProviderGoDaddy {

		godaddyService := service.GodaddyService{
			Config:     conf,
			GodaddyApi: &godaddy.GodaddyApiImpl{Config: conf, HttpClient: httpClient},
			IpObserver: service.IpObserverService{
				IpApi:  &api.IpApiImpl{HttpClient: httpClient},
				Config: conf,
			},
		}

		if err := godaddyService.PrintDomainDetail(); err != nil {
			log.Fatalln(err)
		}

		return &godaddyService, nil

	}

	if conf.Provider == models.ProviderHetzner {
		hetznerService := hetzner.HetznerService{
			Config:     conf,
			HetznerApi: &hetzner.HetznerApiImpl{Config: conf, HttpClient: httpClient},
			IpObserverService: service.IpObserverService{
				IpApi:  &api.IpApiImpl{HttpClient: httpClient},
				Config: conf,
			},
			RecordIds: map[string]hetzner.RecordId{},
		}

		if err := hetznerService.PreloadRecordIds(); err != nil {
			log.Fatalln(err)
		}

		return &hetznerService, nil
	}

	return nil, errors.New("No service for configured provider")
}
