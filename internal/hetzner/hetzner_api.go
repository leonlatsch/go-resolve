package hetzner

import (
	"fmt"

	"github.com/leonlatsch/go-resolve/internal/http"
	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/serialization"
)

type HetznerApi interface {
	GetRecords() ([]Record, error)
	BulkUpdate(records []Record) error
}

type HetznerApiImpl struct {
	Config     *models.Config
	HttpClient http.HttpClient
}

const BASE_URL = "https://dns.hetzner.com/api/v1"

func (api *HetznerApiImpl) GetRecords() ([]Record, error) {
	var records []Record
	var recordsWrapper RecordsWrapper

	url := fmt.Sprintf("%v/records?zone_id=%v", BASE_URL, api.Config.HetznerConfig.ZoneId)
	recordsJson, err := api.HttpClient.Get(url, api.getHeaders())

	if err != nil {
		return records, err
	}

	if err := serialization.FromJson(recordsJson, &recordsWrapper); err != nil {
		return records, err
	}

	return recordsWrapper.Records, nil
}

func (api *HetznerApiImpl) BulkUpdate(records []Record) error {
	url := fmt.Sprintf("%v/records/bulk", BASE_URL)
	recordsWrapper := RecordsWrapper{
		Records: records,
	}
	recordsJson, err := serialization.ToJson(recordsWrapper)
	if err != nil {
		return err
	}

	if _, err := api.HttpClient.Put(url, api.getHeaders(), recordsJson); err != nil {
		return err
	}

	return nil
}

func (api *HetznerApiImpl) getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Auth-API-Token"] = api.Config.HetznerConfig.ApiToken
	return headers
}
