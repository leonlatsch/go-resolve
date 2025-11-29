package hetzner

import (
	"errors"
	"log"

	"github.com/leonlatsch/go-resolve/internal/models"
)

type HetznerService struct {
	Config     *models.Config
	HetznerApi HetznerApi

	RecordIds map[string]RecordId
}

func (service *HetznerService) UpdateDns(ip string) error {
	if len(service.RecordIds) <= 0 {
		return errors.New("no records ids loaded. Not updating")
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

	log.Printf("updating %v records for %v", len(records), service.Config.Domain)
	if err := service.HetznerApi.BulkUpdate(records); err != nil {
		return err
	}

	return nil
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
		return errors.New("could not find configured records in dns entries")
	}

	service.RecordIds = recordIds
	return nil
}
