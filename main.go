package main

import (
	"log"

	"github.com/gin-gonic/gin"
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
		GodaddyApi: &api.GodaddyApiImpl{Config: conf, HttpClient: &httpClient},
		IpApi:      &api.IpApiImpl{HttpClient: &httpClient},
	}

	if err := godaddyService.PrintDomainDetail(); err != nil {
		log.Fatalln(err)
	}
	go godaddyService.ObserveAndUpdateDns()

	serve(&godaddyService)
}

func serve(godaddyService *service.GodaddyService) {
	r := gin.Default()

	r.POST("/api/v1/ip", func(ctx *gin.Context) {
		ip := ctx.Query("newIp")
		godaddyService.OnIpChanged(ip)
	})

	r.Run(":4000")
}
