package service

import (
	"fmt"
	"log"
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

func (self *IpObserverService) ObserveIp(callback func(ip string)) {
	ipChan := self.observePublicIp()

	for {
		ip := <-ipChan
		callback(ip)
	}
}

func (self *IpObserverService) observePublicIp() chan string {
	ipChan := make(chan string)
	interval, err := time.ParseDuration(self.Config.Interval)
	if err != nil {
		log.Println("Could not read retry interval from config. Using fallback 1h")
		interval = time.Hour
	}

	log.Println("Checking for new ip every " + self.Config.Interval)

	go cron.Repeat(interval, func() {
		foundAnyIp := false
		for _, ipApi := range self.Apis {
			currentIp, err := ipApi.GetPublicIpAddress()
			if err != nil {
				log.Println(fmt.Sprintf("Could not get ip from %s:", ipApi.Name()), err)
				continue
			}

			foundAnyIp = true

			if currentIp != self.LastIp {
				log.Printf("Obtained new IP from %s", ipApi.Name())

				ipChan <- currentIp

				// Dont set self.LastIp, main waits for error and handles this
			}
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
