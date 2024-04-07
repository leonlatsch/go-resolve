package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/cron"
	"github.com/leonlatsch/go-resolve/internal/models"
)

type GodaddyService struct {
	Config     models.Config
	GodaddyApi api.GodaddyApi
	IpApi      api.IpApi
	LastIp     string
}

func (self *GodaddyService) PrintDomainDetail() {
	domainDetail, err := self.GodaddyApi.GetDomainDetail()
	if err != nil {
		log.Println("Could not load domain detail for " + self.Config.Domain)
		log.Fatalln(err)
	}

	log.Println(
		fmt.Sprintf(
			"Config valid. Running for domain %s maintained by %s %s (%s)",
			domainDetail.Domain,
			domainDetail.ContactAdmin.FirstName,
			domainDetail.ContactAdmin.LastName,
			domainDetail.ContactAdmin.Email,
		),
	)
}

func (self *GodaddyService) ObserveAndUpdateDns() {
	ipChan := self.observePublicIp()

	for {
		ip := <-ipChan

		if err := self.onIpChanged(ip); err != nil {
			log.Println(err)
		}
	}
}

// Updates all records defined in Hosts with the new ip
func (self *GodaddyService) onIpChanged(ip string) error {
	log.Println("Ip changed: " + ip)
	for _, host := range self.Config.Hosts {

		existingRecords, err := self.GodaddyApi.GetRecords(host)
		if err != nil {
			return err
		}

		record := models.DnsRecord{
			Data: ip,
			Name: host,
			Type: "A",
		}

		switch len(existingRecords) {
		case 0:
			if err := self.GodaddyApi.CreateRecord(record); err != nil {
				return err
			}
		case 1:
			if err := self.GodaddyApi.UpdateRecord(record); err != nil {
				return err
			}
		default:
			return errors.New("Error. Check DNS A records on " + host)
		}
	}

	log.Println("Successfully updated all records. Caching " + ip)
	return nil
}

// Returns a chan string wich will receive a new value once the public ip has changed.
// Uses interval from config or 1h
func (self *GodaddyService) observePublicIp() chan string {
	ipChan := make(chan string)
	interval, err := time.ParseDuration(self.Config.Interval)
	if err != nil {
		log.Println("Could not read retry interval from config. Using fallback 1h")
		interval = time.Hour
	}

	log.Println("Checking for new ip every " + self.Config.Interval)

	go cron.Repeat(interval, func() {
		currentIp, err := self.IpApi.GetPublicIpAddress()
		if err == nil && currentIp != self.LastIp {
			ipChan <- currentIp
			self.LastIp = currentIp
		}
	})

	return ipChan
}
