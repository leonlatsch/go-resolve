package main

import (
	"log"

	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/config"
	"github.com/leonlatsch/go-resolve/internal/http"
	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/service"
)

func main() {
	log.Println("Starting Application")
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	httpClient := http.RealHttpClient{}

	if conf.UpdateUrl == "" {
		startGodaddyMode(&conf, &httpClient)
	} else {
		startUpdateUrlMode(&conf, &httpClient)
	}

}

func startUpdateUrlMode(conf *models.Config, httpClient http.HttpClient) {
	updateUrlService := service.UpdateUrlService{
		Config:       conf,
		UpdateUrlApi: &api.UpdateUrlApiImpl{Config: conf, HttpClient: httpClient},
		IpObserver: service.IpObserver{
			IpApi:  &api.IpApiImpl{HttpClient: httpClient},
			Config: conf,
		},
	}

	updateUrlService.ObserveAndUpdateDns()
}

func startGodaddyMode(conf *models.Config, httpClient http.HttpClient) {
	godaddyService := service.GodaddyService{
		Config:     conf,
		GodaddyApi: &api.GodaddyApiImpl{Config: conf, HttpClient: httpClient},
		IpObserver: service.IpObserver{
			IpApi:  &api.IpApiImpl{HttpClient: httpClient},
			Config: conf,
		},
	}

	if err := godaddyService.PrintDomainDetail(); err != nil {
		log.Fatalln(err)
	}

	godaddyService.ObserveAndUpdateDns()
}
