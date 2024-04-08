package service_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/http"
	"github.com/leonlatsch/go-resolve/internal/models"
	"github.com/leonlatsch/go-resolve/internal/service"
)

func TestPrintDomainDetails(t *testing.T) {
	fakeHttpClient := http.FakeHttpClient{}

	service := service.GodaddyService{
		GodaddyApi: api.GodaddyApi{
			HttpClient: &fakeHttpClient,
		},
	}

	t.Run("Get domain details does not crash with correct json response", func(t *testing.T) {
		fakeDomainDetail := models.DomainDetail{
			Domain: "somedomain.com",
			ContactAdmin: models.DomainContact{
				Email:     "someemail@asdf.com",
				FirstName: "FirstName",
				LastName:  "LastName",
			},
		}

		fakeJson, _ := json.Marshal(fakeDomainDetail)
		fakeHttpClient.RespBody = string(fakeJson)

		if err := service.PrintDomainDetail(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Get domain details crash if http returns an error", func(t *testing.T) {
		fakeHttpClient.Error = errors.New("Some http error")

		if err := service.PrintDomainDetail(); err == nil {
			t.Fatal("Was expected to return error but did not")
		}
	})
}
