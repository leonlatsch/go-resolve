package service

import (
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
		for _, ipApi := range self.Apis {
			currentIp, err := ipApi.GetPublicIpAddress()
			if err == nil && currentIp != self.LastIp {
				ipChan <- currentIp
				// Dont set self.LastIp, main waits for error and handles this
			} else {
				log.Println("Could not get ip from ", err)
			}
		}
	})

	return ipChan
}

func (service *IpObserverService) Initialize() error {
	return nil
}
