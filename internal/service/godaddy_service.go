package service

import (
	"fmt"
	"log"

	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/models"
)

type GodaddyService struct {
	Config     *models.Config
	GodaddyApi api.GodaddyApi
	IpApi      api.IpApi
	IpObserver IpObserver
}

func (self *GodaddyService) PrintDomainDetail() error {
	domainDetail, err := self.GodaddyApi.GetDomainDetail()
	if err != nil {
		log.Println("Could not load domain detail for " + self.Config.Domain)
		return err
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

	return nil
}

func (self *GodaddyService) ObserveAndUpdateDns() {
	self.IpObserver.ObserveIp(func(ip string) {
		self.OnIpChanged(ip)
	})
}

// Updates all records defined in Hosts with the new ip
func (self *GodaddyService) OnIpChanged(ip string) {
	log.Println("Ip changed: " + ip)
	failed := 0

	for _, host := range self.Config.Hosts {

		existingRecords, err := self.GodaddyApi.GetRecords(host)
		if err != nil {
			log.Println(err)
			failed++
			continue
		}

		record := models.DnsRecord{
			Data: ip,
			Name: host,
			Type: "A",
		}

		switch len(existingRecords) {
		case 0:
			if err := self.GodaddyApi.CreateRecord(record); err != nil {
				log.Println(err)
				failed++
				continue
			}
		case 1:
			if err := self.GodaddyApi.UpdateRecord(record); err != nil {
				log.Println(err)
				failed++
				continue
			}
		default:
			log.Println("Error. Check DNS A records on " + host)
			failed++
			continue
		}
	}

	if failed > 0 {
		log.Println("Some updates failed. Not caching ip")
		return
	}

	log.Println("Successfully updated all records. Caching " + ip)
	self.IpObserver.LastIp = ip
}
