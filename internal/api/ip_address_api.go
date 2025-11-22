package api

import (
	"errors"
	"strings"

	"github.com/leonlatsch/go-resolve/internal/http"
)

// Services to resolve the public ipv4 address.
// All of those services return the address as plain text body
var services = []string{
	"https://ipv4.myip.wtf/text",
	"https://api.ipify.org/?format=raw",
	"http://ip.42.pl/raw",
}

type IpApi interface {
	Name() string
	GetPublicIpAddress() (string, error)
}

type IpApiImpl struct {
	HttpClient http.HttpClient
}

func (api *IpApiImpl) Name() string {
	return "External Rest Apis"
}

func (api *IpApiImpl) GetPublicIpAddress() (string, error) {
	for _, service := range services {
		addr, err := api.getIpFrom(service)
		if err == nil && len(addr) != 0 {
			return addr, nil
		}
	}

	return "", errors.New("could not obtain public ip")
}

func (api *IpApiImpl) getIpFrom(service string) (string, error) {
	res, err := api.HttpClient.Get(service, nil)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(res, "\n"), nil
}
