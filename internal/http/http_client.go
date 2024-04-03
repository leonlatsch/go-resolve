package http

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/leonlatsch/go-resolve/internal/serialization"
)

var httpClient = &http.Client{}

func request(method string, url string, headers map[string]string, body interface{}) (string, error) {
    // Serialize req body
	reqBody, err := serialization.ToJson(body)
	if err != nil {
		return "", err
	}

    // Create request
	req, err := http.NewRequest(method, url, bytes.NewBufferString(reqBody))
	if err != nil {
		return "", err
	}

    // Add headers
    req.Header.Add("Content-Type", "application/json")
    for key, value := range headers {
        req.Header.Add(key, value)
    } 

    // Execute request
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

    // Read resp body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

    respOk := resp.StatusCode >= 200 && resp.StatusCode < 300
    if !respOk {
        return string(respBody), errors.New(resp.Status)
    }

    return string(respBody), nil
}

func GET(url string, headers map[string]string) (string, error) {
    respBody, err := request("GET", url, headers, nil)
    if err != nil {
        return "", err
    }

    return respBody, nil
}

func PUT(url string, headers map[string]string, requestBody interface{}) (string, error) {
    respBody, err := request("PUT", url, headers,  requestBody)
    if err != nil {
        return "", err
    }

    return respBody, nil
}


func PATCH(url string, headers map[string]string, requestBody interface{}) (string, error) {
    respBody, err := request("PATCH", url, headers,  requestBody)
    if err != nil {
        return "", err
    }

    return respBody, nil
}
