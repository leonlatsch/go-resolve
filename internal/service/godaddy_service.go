package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/leonlatsch/go-resolve/internal/godaddy"
	"github.com/leonlatsch/go-resolve/internal/models"
)

type GodaddyService struct {
	Config     *models.Config
	GodaddyApi godaddy.GodaddyApi
}

// Updates all records defined in Hosts with the new ip
func (service *GodaddyService) UpdateDns(ip string) error {
	failed := 0

	for _, host := range service.Config.Hosts {

		existingRecords, err := service.GodaddyApi.GetRecords(host)
		if err != nil {
			log.Println(err)
			failed++
			continue
		}

		record := godaddy.DnsRecord{
			Data: ip,
			Name: host,
			Type: "A",
		}

		switch len(existingRecords) {
		case 0:
			if err := service.GodaddyApi.CreateRecord(record); err != nil {
				log.Println(err)
				failed++
				continue
			}
		case 1:
			if err := service.GodaddyApi.UpdateRecord(record); err != nil {
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
		return errors.New("Could not update all records")
	}

	return nil
}

func (service *GodaddyService) Initialize() error {
	domainDetail, err := service.GodaddyApi.GetDomainDetail()
	if err != nil {
		log.Println("Could not load domain detail for " + service.Config.Domain)
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
