package api

import (
	"fmt"
	"log"

	"github.com/leonlatsch/go-resolve/config"
	"github.com/leonlatsch/go-resolve/http"
	"github.com/leonlatsch/go-resolve/models"
	"github.com/leonlatsch/go-resolve/serialization"
)

const BASE_URL = "https://api.godaddy.com/v1"

func endpointRecords(host string) string {
    return BASE_URL + fmt.Sprintf("/domains/%s/records/%s", config.SharedConfig.Domain, host)
}

func endpointARecords(host string) string {
    return BASE_URL + fmt.Sprintf("/domains/%s/records/A/%s", config.SharedConfig.Domain, host)
}

func endpointDomainDetail() string {
    return BASE_URL + fmt.Sprintf("/domains/%s", config.SharedConfig.Domain)
}

func getAuthHeaders() map[string]string {
    headers := make(map[string]string)
    headers["Authorization"] = "sso-key " + config.SharedConfig.ApiKey + ":" + config.SharedConfig.ApiSecret
    return headers
}

func GetDomainDetail() (models.DomainDetail, error) {
    var detail models.DomainDetail
    json, err := http.GET("", getAuthHeaders())
    if err != nil {
        return detail, err
    }

    if err := serialization.FromJson(json, &detail); err != nil {
        return detail, err
    }

    return detail, nil
}

func GetRecords(host string) ([]models.DnsRecord, error) {
    var records []models.DnsRecord
    recordsJson, err := http.GET(endpointARecords(host), getAuthHeaders())

    if err != nil {
        return records, err
    }
    
    if err := serialization.FromJson(recordsJson, &records); err != nil {
        return records, err
    } 

    return records, nil
} 

func CreateRecord(record models.DnsRecord) error {
    log.Println("Creating " + record.Name + "." + config.SharedConfig.Domain + " -> " + record.Data)
    records := []models.DnsRecord{record}

    if _, err := http.PATCH(endpointRecords(""), getAuthHeaders(), records); err != nil {
        return err
    }

    return nil
}

func UpdateRecord(record models.DnsRecord) error {
    log.Println("Updating " + record.Name + "." + config.SharedConfig.Domain + " -> " + record.Data)
    records := []models.DnsRecord{record}

    if _, err := http.PUT(endpointARecords(record.Name), getAuthHeaders(), records); err != nil {
        return err
    }

    return nil
}

