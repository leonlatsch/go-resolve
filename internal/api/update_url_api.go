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

func (self *UpdateUrlApiImpl) CallUpdateUrl(host string) error {
	url := strings.ReplaceAll(self.Config.UpdateUrl, "<host>", host)
	url = strings.ReplaceAll(url, "<domain>", self.Config.Domain)
	_, err := self.HttpClient.Get(url, nil)
	if err != nil {
		return err
	}

	return nil

}
