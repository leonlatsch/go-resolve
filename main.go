package main

import (
	"log"

	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/config"
	"github.com/leonlatsch/go-resolve/internal/http"
	"github.com/leonlatsch/go-resolve/internal/service"
)

func main() {
	log.Println("Starting Application")
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	httpClient := http.RealHttpClient{}
	godaddyService := service.GodaddyService{
		Config:     conf,
		GodaddyApi: api.GodaddyApi{Config: conf, HttpClient: &httpClient},
		IpApi:      api.IpApi{HttpClient: &httpClient},
	}

	godaddyService.PrintDomainDetail()
	godaddyService.ObserveAndUpdateDns()
}
