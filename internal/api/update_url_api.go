package api

import (
	"strings"

	"github.com/leonlatsch/go-resolve/internal/http"
	"github.com/leonlatsch/go-resolve/internal/models"
)

type UpdateUrlApi interface {
	CallUpdateUrl(host string) error
}

type UpdateUrlApiImpl struct {
	Config     *models.Config
	HttpClient http.HttpClient
}

func (api *UpdateUrlApiImpl) CallUpdateUrl(host string) error {
	url := strings.ReplaceAll(api.Config.UpdateUrlConfig.Url, "<host>", host)
	url = strings.ReplaceAll(url, "<domain>", api.Config.Domain)
	_, err := api.HttpClient.Get(url, nil)
	if err != nil {
		return err
	}

	return nil
}
