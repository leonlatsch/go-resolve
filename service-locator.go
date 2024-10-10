package main

import (
	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/godaddy"
	"github.com/leonlatsch/go-resolve/internal/hetzner"
	"github.com/leonlatsch/go-resolve/internal/http"
	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/service"
)

var ServiceLocator *serviceLocator

type serviceLocator struct {
	HttpClient http.HttpClient

	IpApi        api.IpApi
	UpdateUrlApi api.UpdateUrlApi
	GoDaddyApi   godaddy.GodaddyApi
	HetznerApi   hetzner.HetznerApi

	IpObserverService service.IpObserverService
	UpdateUrlService  service.UpdateUrlService
	GoDaddyService    service.GodaddyService
	HetznerService    hetzner.HetznerService
}

func InitializeServiceLocator(conf *models.Config) {
	httpClient := &http.RealHttpClient{}

	ipApi := &api.IpApiImpl{
		HttpClient: httpClient,
	}

	updateUrlApi := &api.UpdateUrlApiImpl{
		Config:     conf,
		HttpClient: httpClient,
	}

	godaddyApi := &godaddy.GodaddyApiImpl{
		Config:     conf,
		HttpClient: httpClient,
	}

	hetznerApi := &hetzner.HetznerApiImpl{
		Config:     conf,
		HttpClient: httpClient,
	}

	ipObserverService := service.IpObserverService{
		IpApi:  ipApi,
		Config: conf,
	}

	updateUrlService := service.UpdateUrlService{
		Config:       conf,
		UpdateUrlApi: updateUrlApi,
		IpObserver:   ipObserverService,
	}

	godaddyService := service.GodaddyService{
		Config:     conf,
		GodaddyApi: godaddyApi,
		IpApi:      ipApi,
		IpObserver: ipObserverService,
	}

	hetznerService := hetzner.HetznerService{
		Config:            conf,
		HetznerApi:        hetznerApi,
		IpObserverService: ipObserverService,
	}

	ServiceLocator = &serviceLocator{
		HttpClient: httpClient,

		IpApi:        ipApi,
		UpdateUrlApi: updateUrlApi,
		GoDaddyApi:   godaddyApi,
		HetznerApi:   hetznerApi,

		IpObserverService: ipObserverService,
		UpdateUrlService:  updateUrlService,
		GoDaddyService:    godaddyService,
		HetznerService:    hetznerService,
	}
}
