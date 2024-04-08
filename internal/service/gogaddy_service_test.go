package service_test

import (
	"errors"
	"testing"

	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/service"
)

func TestPrintDomainDetails(t *testing.T) {
	godaddyApiFake := api.GodaddyApiFake{}

	service := service.GodaddyService{
		GodaddyApi: &godaddyApiFake,
		IpApi:      &api.IpApiFake{},
	}

	t.Run("Get domain details does not crash with correct json response", func(t *testing.T) {
		fakeDomainDetail := models.DomainDetail{
			Domain: "somedomain.com",
			ContactAdmin: models.DomainContact{
				Email:     "someemail@asdf.com",
				FirstName: "FirstName",
				LastName:  "LastName",
			},
		}

		godaddyApiFake.DomainDetail = fakeDomainDetail

		if err := service.PrintDomainDetail(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Get domain details crash if http returns an error", func(t *testing.T) {
		godaddyApiFake.Error = errors.New("Some http error")

		if err := service.PrintDomainDetail(); err == nil {
			t.Fatal("Was expected to return error but did not")
		}
	})
}

func TestObserveAndUpdateDns(t *testing.T) {

	godaddyApiFake := api.GodaddyApiFake{}

	service := service.GodaddyService{
		GodaddyApi: &godaddyApiFake,
		IpApi:      &api.IpApiFake{},
	}

	fakeRecords := []models.DnsRecord{
		{
			Data: "someIp",
			Name: "someName",
			Type: "A",
		},
	}

	godaddyApiFake.ExistingRecords = fakeRecords
	godaddyApiFake.Error = nil
}
