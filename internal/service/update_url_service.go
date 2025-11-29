package service

import (
	"errors"
	"log"

	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/models"
)

type UpdateUrlService struct {
	Config       *models.Config
	UpdateUrlApi api.UpdateUrlApi
}

func (service *UpdateUrlService) UpdateDns(ip string) error {
	failed := 0
	for _, host := range service.Config.Hosts {
		if err := service.UpdateUrlApi.CallUpdateUrl(host); err != nil {
			log.Println("Could not update via url for host " + host)
			failed++
			continue
		}

		log.Println("Updating " + host + "." + service.Config.Domain)
	}

	if failed > 0 {
		return errors.New("could not update all records")
	}

	return nil
}

func (service *UpdateUrlService) Initialize() error {
	return nil
}
