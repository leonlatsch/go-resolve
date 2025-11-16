package hetznercloud

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	"github.com/leonlatsch/go-resolve/internal/api"
	"github.com/leonlatsch/go-resolve/internal/models"
)

type HetznerCloudService struct {
	Config *models.Config
	IpApi  api.IpApi

	client *hcloud.Client
	zone   *hcloud.Zone
}

func (service *HetznerCloudService) UpdateDns(ip string) error {
	failed := 0
	for _, host := range service.Config.Hosts {
		log.Println("Updating records for " + host + "." + service.Config.Domain)
		rrset, _, err := service.client.Zone.GetRRSetByNameAndType(context.Background(), service.zone, host, hcloud.ZoneRRSetTypeA)
		if err != nil {
			log.Println("No RRSet for host "+host+". Skipping", err)
			failed++
			continue
		}

		setOpts := hcloud.ZoneRRSetSetRecordsOpts{
			Records: []hcloud.ZoneRRSetRecord{
				{Value: ip, Comment: "updated via go-resolve"},
			},
		}

		action, _, err := service.client.Zone.SetRRSetRecords(context.Background(), rrset, setOpts)
		if err != nil {
			log.Println("Error updating records for RRSet", err)
			failed++
			continue
		}

		if action != nil {
			waitCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			if err := service.client.Action.WaitFor(waitCtx, action); err != nil {
				log.Println("Error while waiting for timeout", err)
				cancel()
				failed++
				continue
			}
			cancel()
		}
	}

	if failed > 0 {
		return errors.New("Cloud not update all records")
	}

	return nil
}

func (service *HetznerCloudService) Initialize() error {
	service.client = hcloud.NewClient(hcloud.WithToken(service.Config.HetznerCloudConfig.CloudApiToken))

	zone, _, err := service.client.Zone.GetByName(context.Background(), service.Config.Domain)
	if err != nil {
		log.Println("Error getting zone", err)
		return err
	}

	service.zone = zone

	existingsRrSets, _, err := service.client.Zone.ListRRSets(
		context.Background(),
		zone, hcloud.ZoneRRSetListOpts{
			Type: []hcloud.ZoneRRSetType{hcloud.ZoneRRSetTypeA},
		},
	)
	if err != nil {
		return err
	}

	currentIP, err := service.IpApi.GetPublicIpAddress()
	if err != nil {
		return err
	}

	for _, host := range service.Config.Hosts {
		var exsits bool
		for _, rrset := range existingsRrSets {
			if rrset.Name == host {
				exsits = true
				break
			}
		}

		if !exsits {
			log.Println("Host " + host + " does not exist in zone " + zone.Name + ". Creating")

			_, _, err := service.client.Zone.CreateRRSet(
				context.Background(),
				zone, hcloud.ZoneRRSetCreateOpts{
					Name: host,
					Type: hcloud.ZoneRRSetTypeA,
					TTL:  hcloud.Ptr(3600),
					Records: []hcloud.ZoneRRSetRecord{
						{
							Value:   currentIP,
							Comment: "updated via go-resolve",
						},
					},
				},
			)
			if err != nil {
				log.Println("Error creating RR set for " + host + ". Fix your config or create the record yourself in hetzner cloud console. Exiting.")
				return err
			}
		}
	}

	return nil
}
