package api

import (
	"fmt"
	"log"

	"github.com/leonlatsch/go-resolve/internal/http"
	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/serialization"
)

type GodaddyApi struct {
	Config models.Config
}

const BASE_URL = "https://api.godaddy.com/v1"

func (self *GodaddyApi) GetDomainDetail() (models.DomainDetail, error) {
	var detail models.DomainDetail
	json, err := http.GET(self.endpointDomainDetail(), self.getAuthHeaders())
	if err != nil {
		return detail, err
	}

	if err := serialization.FromJson(json, &detail); err != nil {
		return detail, err
	}

	return detail, nil
}

func (self *GodaddyApi) GetRecords(host string) ([]models.DnsRecord, error) {
	var records []models.DnsRecord
	recordsJson, err := http.GET(self.endpointARecords(host), self.getAuthHeaders())

	if err != nil {
		return records, err
	}

	if err := serialization.FromJson(recordsJson, &records); err != nil {
		return records, err
	}

	return records, nil
}

func (self *GodaddyApi) CreateRecord(record models.DnsRecord) error {
	log.Println("Creating " + record.Name + "." + self.Config.Domain + " -> " + record.Data)
	records := []models.DnsRecord{record}

	if _, err := http.PATCH(self.endpointRecords(""), self.getAuthHeaders(), records); err != nil {
		return err
	}

	return nil
}

func (self *GodaddyApi) UpdateRecord(record models.DnsRecord) error {
	log.Println("Updating " + record.Name + "." + self.Config.Domain + " -> " + record.Data)
	records := []models.DnsRecord{record}

	if _, err := http.PUT(self.endpointARecords(record.Name), self.getAuthHeaders(), records); err != nil {
		return err
	}

	return nil
}

func (self *GodaddyApi) endpointRecords(host string) string {
	return BASE_URL + fmt.Sprintf("/domains/%s/records/%s", self.Config.Domain, host)
}

func (self *GodaddyApi) endpointARecords(host string) string {
	return BASE_URL + fmt.Sprintf("/domains/%s/records/A/%s", self.Config.Domain, host)
}

func (self *GodaddyApi) endpointDomainDetail() string {
	return BASE_URL + fmt.Sprintf("/domains/%s", self.Config.Domain)
}

func (self *GodaddyApi) getAuthHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = "sso-key " + self.Config.ApiKey + ":" + self.Config.ApiSecret
	return headers
}
