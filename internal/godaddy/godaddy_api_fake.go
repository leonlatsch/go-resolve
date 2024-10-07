package godaddy

type GodaddyApiFake struct {
	DomainDetail    DomainDetail
	ExistingRecords map[string][]DnsRecord
	Error           error

	CreateRecordCalledWith DnsRecord
	UpdateRecordCalledWith DnsRecord
}

func (self *GodaddyApiFake) GetDomainDetail() (DomainDetail, error) {
	return self.DomainDetail, self.Error
}

func (self *GodaddyApiFake) GetRecords(host string) ([]DnsRecord, error) {
	return self.ExistingRecords[host], self.Error
}

func (self *GodaddyApiFake) CreateRecord(record DnsRecord) error {
	self.CreateRecordCalledWith = record
	return self.Error
}

func (self *GodaddyApiFake) UpdateRecord(record DnsRecord) error {
	self.UpdateRecordCalledWith = record
	return self.Error
}
