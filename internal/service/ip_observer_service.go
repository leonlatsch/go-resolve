package service

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/cron"
	"github.com/leonlatsch/go-resolve/internal/models"
)

type IpObserverService struct {
	Apis   []api.IpApi
	Config *models.Config
	LastIp string
}

func (service *IpObserverService) PrintIpProviders() {
	names := []string{}

	for _, api := range service.Apis {
		names = append(names, api.Name())
	}

	log.Printf("Configured IP Providers: %s", strings.Join(names, ", "))
}

func (service *IpObserverService) ObserveIp(callback func(ip string)) {
	ipChan := service.observePublicIp()

	for {
		ip := <-ipChan
		callback(ip)
	}
}

func (service *IpObserverService) observePublicIp() chan string {
	ipChan := make(chan string)
	interval, err := time.ParseDuration(service.Config.Interval)
	if err != nil {
		log.Println("Could not read retry interval from config. Using fallback 1h")
		interval = time.Hour
	}

	log.Println("Checking for new ip every " + service.Config.Interval)

	go cron.Repeat(interval, func() {
		foundAnyIp := false
		for _, ipApi := range service.Apis {
			currentIp, err := ipApi.GetPublicIpAddress()
			if err != nil {
				log.Println(fmt.Sprintf("Could not get ip from %s:", ipApi.Name()), err)
				continue
			}

			if currentIp != service.LastIp {
				ipChan <- currentIp
			}

			// Found a IP with this provider. Break the loop and dont try other providers
			foundAnyIp = true
			break
		}

		if !foundAnyIp {
			log.Println("Could not obtain a IP from any provider. Skipping update.")
		}
	})

	return ipChan
}

func (service *IpObserverService) Initialize() error {
	return nil
}
