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
	ipOpserver := service.IpObserverService{}

	service := service.GodaddyService{
		Config:     &models.Config{},
		GodaddyApi: &godaddyApiFake,
		IpApi:      &api.IpApiFake{},
		IpObserver: ipOpserver,
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

func TestOnIpChanged(t *testing.T) {
	godaddyApiFake := api.GodaddyApiFake{}
	conf := models.Config{
		Domain: "mydomain.com",
		Hosts:  []string{"host1", "host2"},
	}

	service := service.GodaddyService{
		Config:     &conf,
		GodaddyApi: &godaddyApiFake,
		IpApi:      &api.IpApiFake{},
	}

	t.Run("OnIpChanged creates new and updates existing record", func(t *testing.T) {
		godaddyApiFake.ExistingRecords = make(map[string][]models.DnsRecord)
		godaddyApiFake.ExistingRecords["host1"] = []models.DnsRecord{
			{
				Data: "oldIp",
				Name: "host1",
				Type: "A",
			},
		}
		godaddyApiFake.ExistingRecords["host2"] = []models.DnsRecord{}

		// host1 should be updated and host2 should be created

		newIp := "123.123.123.123"
		service.UpdateDns(newIp)

		expectedUpdatedRecord := models.DnsRecord{
			Data: newIp,
			Name: "host1",
			Type: "A",
		}
		expectedCreatedRecord := models.DnsRecord{
			Data: newIp,
			Name: "host2",
			Type: "A",
		}

		if godaddyApiFake.UpdateRecordCalledWith != expectedUpdatedRecord {
			t.Error("Updated record did not match expected")
		}
		if godaddyApiFake.CreateRecordCalledWith != expectedCreatedRecord {
			t.Error("Created record did not match expected")
		}
		if service.IpObserver.LastIp != newIp {
			t.Error("New ip was not cached")
		}
	})
}
