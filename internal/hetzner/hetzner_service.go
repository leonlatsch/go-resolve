package hetzner

import (
	"errors"
	"fmt"
	"log"

	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/service"
)

type HetznerService struct {
	Config            *models.Config
	HetznerApi        HetznerApi
	IpObserverService service.IpObserverService

	RecordIds map[string]RecordId
}

func (service *HetznerService) ObserveAndUpdateDns() {
	log.Println("Running for hetzner")
	service.IpObserverService.ObserveIp(func(ip string) {
		service.UpdateDns(ip)
	})
}

func (service *HetznerService) UpdateDns(ip string) {
	log.Println("Ip changed: " + ip)

	if len(service.RecordIds) <= 0 {
		log.Println("No records ids loaded. Not updating.")
		return
	}

	records := []Record{}
	for _, host := range service.Config.Hosts {
		record := Record{
			Id:     service.RecordIds[host],
			Value:  ip,
			Type:   "A",
			Name:   host,
			ZoneId: ZoneId(service.Config.HetznerConfig.ZoneId),
		}
		records = append(records, record)
	}

	log.Println(fmt.Sprintf("Updating %v records for %v", len(records), service.Config.Domain))
	if err := service.HetznerApi.BulkUpdate(records); err != nil {
		log.Println("Bulk update failed. Not caching ip")
		log.Println(err)
		return
	}

	log.Println("Successfully updated all records. Caching " + ip)
	service.IpObserverService.LastIp = ip
}

func (service *HetznerService) Initialize() error {
	recordIds := make(map[string]RecordId)

	records, err := service.HetznerApi.GetRecords()
	if err != nil {
		log.Println("Could not preload records ids. Please check your config")
		return err
	}

	for _, record := range records {
		for _, host := range service.Config.Hosts {
			if host == record.Name {
				log.Printf("Loaded record id %v for %v", record.Id, host)
				recordIds[host] = record.Id
			}
		}
	}

	if len(recordIds) <= 0 {
		return errors.New("Could not find configured records in dns entries")
	}

	service.RecordIds = recordIds
	return nil
}
