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
	GetPublicIpAddress() (string, error)
}

type IpApiImpl struct {
	HttpClient http.HttpClient
}

func (self *IpApiImpl) GetPublicIpAddress() (string, error) {
	for _, service := range services {
		addr, err := self.getIpFrom(service)
		if err == nil && len(addr) != 0 {
			return addr, nil
		}
	}

	return "", errors.New("Could not obtain public ip")
}

func (self *IpApiImpl) getIpFrom(service string) (string, error) {
	res, err := self.HttpClient.Get(service, nil)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(res, "\n"), nil
}
