package service

import (
	"log"

	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/models"
)

type UpdateUrlService struct {
	Config       *models.Config
	UpdateUrlApi api.UpdateUrlApi
	IpObserver   IpObserver
}

func (self *UpdateUrlService) ObserveAndUpdateDns() {
    self.IpObserver.ObserveIp(func(ip string) {
        self.OnIpChanged()
    })
}

func (self *UpdateUrlService) OnIpChanged() {
	for _, host := range self.Config.Hosts {
		if err := self.UpdateUrlApi.CallUpdateUrl(host); err != nil {
			log.Println("Could not update via url for host " + host)
			continue
		}

		log.Println("Updating " + host + "." + self.Config.Domain)
	}
}
