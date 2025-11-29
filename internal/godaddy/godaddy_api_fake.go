package godaddy

type GodaddyApiFake struct {
	DomainDetail    DomainDetail
	ExistingRecords map[string][]DnsRecord
	Error           error

	CreateRecordCalledWith DnsRecord
	UpdateRecordCalledWith DnsRecord
}

func (fake *GodaddyApiFake) GetDomainDetail() (DomainDetail, error) {
	return fake.DomainDetail, fake.Error
}

func (fake *GodaddyApiFake) GetRecords(host string) ([]DnsRecord, error) {
	return fake.ExistingRecords[host], fake.Error
}

func (fake *GodaddyApiFake) CreateRecord(record DnsRecord) error {
	fake.CreateRecordCalledWith = record
	return fake.Error
}

func (fake *GodaddyApiFake) UpdateRecord(record DnsRecord) error {
	fake.UpdateRecordCalledWith = record
	return fake.Error
}
