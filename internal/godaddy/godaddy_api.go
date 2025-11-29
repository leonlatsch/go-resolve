package godaddy

import (
	"fmt"
	"log"

	"github.com/leonlatsch/go-resolve/internal/http"
	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/serialization"
)

type GodaddyApi interface {
	GetDomainDetail() (DomainDetail, error)
	GetRecords(host string) ([]DnsRecord, error)
	CreateRecord(record DnsRecord) error
	UpdateRecord(record DnsRecord) error
}

type GodaddyApiImpl struct {
	Config     *models.Config
	HttpClient http.HttpClient
}

const BASE_URL = "https://api.godaddy.com/v1"

func (api *GodaddyApiImpl) GetDomainDetail() (DomainDetail, error) {
	var detail DomainDetail
	json, err := api.HttpClient.Get(api.endpointDomainDetail(), api.getAuthHeaders())
	if err != nil {
		return detail, err
	}

	if err := serialization.FromJson(json, &detail); err != nil {
		return detail, err
	}

	return detail, nil
}

func (api *GodaddyApiImpl) GetRecords(host string) ([]DnsRecord, error) {
	var records []DnsRecord
	recordsJson, err := api.HttpClient.Get(api.endpointARecords(host), api.getAuthHeaders())
	if err != nil {
		return records, err
	}

	if err := serialization.FromJson(recordsJson, &records); err != nil {
		return records, err
	}

	return records, nil
}

func (api *GodaddyApiImpl) CreateRecord(record DnsRecord) error {
	log.Println("Creating " + record.Name + "." + api.Config.Domain + " -> " + record.Data)
	records := []DnsRecord{record}

	if _, err := api.HttpClient.Patch(api.endpointRecords(""), api.getAuthHeaders(), records); err != nil {
		return err
	}

	return nil
}

func (api *GodaddyApiImpl) UpdateRecord(record DnsRecord) error {
	log.Println("Updating " + record.Name + "." + api.Config.Domain + " -> " + record.Data)
	records := []DnsRecord{record}

	if _, err := api.HttpClient.Put(api.endpointARecords(record.Name), api.getAuthHeaders(), records); err != nil {
		return err
	}

	return nil
}

func (api *GodaddyApiImpl) endpointRecords(host string) string {
	return BASE_URL + fmt.Sprintf("/domains/%s/records/%s", api.Config.Domain, host)
}

func (api *GodaddyApiImpl) endpointARecords(host string) string {
	return BASE_URL + fmt.Sprintf("/domains/%s/records/A/%s", api.Config.Domain, host)
}

func (api *GodaddyApiImpl) endpointDomainDetail() string {
	return BASE_URL + fmt.Sprintf("/domains/%s", api.Config.Domain)
}

func (api *GodaddyApiImpl) getAuthHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = "sso-key " + api.Config.GoDaddyConfig.ApiKey + ":" + api.Config.GoDaddyConfig.ApiSecret
	return headers
}
