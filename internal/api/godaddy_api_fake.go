package api

import "github.com/leonlatsch/go-resolve/internal/models"

type GodaddyApiFake struct {
	DomainDetail    models.DomainDetail
	ExistingRecords []models.DnsRecord
	Error           error

	CreateRecordCalledWith models.DnsRecord
	UpdateRecordCalledWith models.DnsRecord
}

func (self *GodaddyApiFake) GetDomainDetail() (models.DomainDetail, error) {
	return self.DomainDetail, self.Error
}

func (self *GodaddyApiFake) GetRecords(host string) ([]models.DnsRecord, error) {
	return self.ExistingRecords, self.Error
}

func (self *GodaddyApiFake) CreateRecord(record models.DnsRecord) error {
	self.CreateRecordCalledWith = record
	return self.Error
}

func (self *GodaddyApiFake) UpdateRecord(record models.DnsRecord) error {
	self.UpdateRecordCalledWith = record
	return self.Error
}
