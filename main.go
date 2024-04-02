package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/leonlatsch/go-resolve/api"
	"github.com/leonlatsch/go-resolve/config"
	"github.com/leonlatsch/go-resolve/cron"
	"github.com/leonlatsch/go-resolve/models"
)

func main() {
	log.Println("Starting Application")
	config.LoadConfig()

    domainDetail, err := api.GetDomainDetail()
    if err != nil {
        log.Println("Could not load domain detail for " + config.SharedConfig.Domain)
        log.Fatalln(err)
    }

    log.Println(
        fmt.Sprintf(
            "Config valid. Running for domain %s maintained by %s %s (%s)",
            domainDetail.Domain,
            domainDetail.ContactAdmin.FirstName,
            domainDetail.ContactAdmin.LastName,
            domainDetail.ContactAdmin.Email,
        ),
    )

	ipChan := observePublicIp()
	lastIp := ""

	for {
		ip := <-ipChan

		if ip == lastIp {
			continue
		}

		if err := onIpChanged(ip); err == nil {
			log.Println("Successfully updated all records. Caching " + ip)
			lastIp = ip
		} else {
			log.Println(err)
		}
	}
}

func onIpChanged(ip string) error {
    log.Println("Ip changed: " + ip)
	for _, host := range config.SharedConfig.Hosts {

		existingRecords, err := api.GetRecords(host)
		if err != nil {
			return err
		}

		record := models.DnsRecord{
			Data: ip,
			Name: host,
			Type: "A",
		}

		switch len(existingRecords) {
		case 0:
			if err := api.CreateRecord(record); err != nil {
				return err
			}
		case 1:
			if err := api.UpdateRecord(record); err != nil {
				return err
			}
		default:
			return errors.New("Error. Check DNS A records on " + host)
		}
	}
	return nil
}

func observePublicIp() chan string {
	ipChan := make(chan string)
	interval, err := time.ParseDuration(config.SharedConfig.Interval)
	if err != nil {
		log.Println("Could not read retry interval from config. Using fallback 1h")
		interval = time.Hour
	}

    log.Println("Checking for new ip every " + config.SharedConfig.Interval)
	go cron.RunAndRepeat(interval, func() {
		currentIp, err := api.GetPublicIpAddress()
		if err == nil {
			ipChan <- currentIp
		}
	})

	return ipChan
}
