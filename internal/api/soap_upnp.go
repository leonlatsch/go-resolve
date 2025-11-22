package api

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"regexp"
)

type UpnpIPAPI struct{}

var reqBody = `<?xml version="1.0" encoding="utf-8"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
  <s:Body>
    <u:GetExternalIPAddress xmlns:u="urn:schemas-upnp-org:service:WANIPConnection:1" />
  </s:Body>
</s:Envelope>`

func (api *UpnpIPAPI) GetPublicIpAddress() (string, error) {
	req, _ := http.NewRequest("POST",
		"http://10.10.0.1:49000/igdupnp/control/WANIPConn1",
		bytes.NewBufferString(reqBody))

	req.Header.Set("Content-Type", `text/xml; charset="utf-8"`)
	req.Header.Set("SOAPAction",
		`urn:schemas-upnp-org:service:WANIPConnection:1#GetExternalIPAddress`)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)

	re := regexp.MustCompile(`<NewExternalIPAddress>([^<]+)</NewExternalIPAddress>`)
	match := re.FindStringSubmatch(string(resBody))
	if len(match) < 2 {
		return "", errors.New("no IP found at gateway")
	}

	return match[1], nil
}
