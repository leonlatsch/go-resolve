package main

import (
	"log"

	"github.com/leonlatsch/go-resolve/internal/config"
	"github.com/leonlatsch/go-resolve/internal/service"
)

func main() {
	log.Println("Starting Application")
	config.LoadConfig()

    godaddyService := service.GodaddyService{Config: config.SharedConfig}
    godaddyService.PrintDomainDetail()

    godaddyService.ObserveAndUpdateDns()
}

