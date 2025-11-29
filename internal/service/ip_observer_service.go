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

func (service *IpObserverService) ObserveIpAndNotify(dnsProvider DnsModeService) {
	interval, err := time.ParseDuration(service.Config.Interval)
	if err != nil {
		log.Println("Could not read retry interval from config. Using fallback 1h")
		interval = time.Hour
	}

	log.Println("Checking for new ip every " + service.Config.Interval)

	cron.Repeat(interval, func() {
		foundAnyIp := false
		for _, ipApi := range service.Apis {
			currentIp, err := ipApi.GetPublicIpAddress()
			if err != nil {
				log.Println(fmt.Sprintf("Could not get ip from %s:", ipApi.Name()), err)
				continue
			}

			foundAnyIp = true

			// If up is different callback is exec
			if currentIp != service.LastIp {
				if err := dnsProvider.UpdateDns(currentIp); err == nil {
					service.LastIp = currentIp
					log.Println("Successfully updated all records. Caching " + currentIp)
				} else {
					log.Println("Not caching ip: ", err)
				}
			}

			// Found a IP with this provider. Break the loop and dont try other providers
			break
		}

		if !foundAnyIp {
			log.Println("Could not obtain a IP from any provider. Skipping update.")
		}
	})
}

func (service *IpObserverService) Initialize() error {
	return nil
}
