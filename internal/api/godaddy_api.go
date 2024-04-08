package api

import (
	"fmt"
	"log"

	"github.com/leonlatsch/go-resolve/internal/http"
	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/serialization"
)

type GodaddyApi interface {
	GetDomainDetail() (models.DomainDetail, error)
	GetRecords(host string) ([]models.DnsRecord, error)
	CreateRecord(record models.DnsRecord) error
	UpdateRecord(record models.DnsRecord) error
}

type GodaddyApiImpl struct {
	Config     models.Config
	HttpClient http.HttpClient
}

const BASE_URL = "https://api.godaddy.com/v1"

func (self *GodaddyApiImpl) GetDomainDetail() (models.DomainDetail, error) {
	var detail models.DomainDetail
	json, err := self.HttpClient.Get(self.endpointDomainDetail(), self.getAuthHeaders())
	if err != nil {
		return detail, err
	}

	if err := serialization.FromJson(json, &detail); err != nil {
		return detail, err
	}

	return detail, nil
}

func (self *GodaddyApiImpl) GetRecords(host string) ([]models.DnsRecord, error) {
	var records []models.DnsRecord
	recordsJson, err := self.HttpClient.Get(self.endpointARecords(host), self.getAuthHeaders())

	if err != nil {
		return records, err
	}

	if err := serialization.FromJson(recordsJson, &records); err != nil {
		return records, err
	}

	return records, nil
}

func (self *GodaddyApiImpl) CreateRecord(record models.DnsRecord) error {
	log.Println("Creating " + record.Name + "." + self.Config.Domain + " -> " + record.Data)
	records := []models.DnsRecord{record}

	if _, err := self.HttpClient.Patch(self.endpointRecords(""), self.getAuthHeaders(), records); err != nil {
		return err
	}

	return nil
}

func (self *GodaddyApiImpl) UpdateRecord(record models.DnsRecord) error {
	log.Println("Updating " + record.Name + "." + self.Config.Domain + " -> " + record.Data)
	records := []models.DnsRecord{record}

	if _, err := self.HttpClient.Put(self.endpointARecords(record.Name), self.getAuthHeaders(), records); err != nil {
		return err
	}

	return nil
}

func (self *GodaddyApiImpl) endpointRecords(host string) string {
	return BASE_URL + fmt.Sprintf("/domains/%s/records/%s", self.Config.Domain, host)
}

func (self *GodaddyApiImpl) endpointARecords(host string) string {
	return BASE_URL + fmt.Sprintf("/domains/%s/records/A/%s", self.Config.Domain, host)
}

func (self *GodaddyApiImpl) endpointDomainDetail() string {
	return BASE_URL + fmt.Sprintf("/domains/%s", self.Config.Domain)
}

func (self *GodaddyApiImpl) getAuthHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = "sso-key " + self.Config.ApiKey + ":" + self.Config.ApiSecret
	return headers
}
