package main

import (
	"log"

	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/config"
	"github.com/leonlatsch/go-resolve/internal/service"
)

func main() {
	log.Println("Starting Application")
    conf, err:= config.LoadConfig()
    if err != nil {
        log.Fatalln(err)
    }

	godaddyService := service.GodaddyService{
		Config:     conf,
		GodaddyApi: api.GodaddyApi{Config: conf},
	}

	godaddyService.PrintDomainDetail()
	godaddyService.ObserveAndUpdateDns()
}
