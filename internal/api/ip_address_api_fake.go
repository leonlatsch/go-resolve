package api

type IpApiFake struct {
	Ip    string
	Error error
}

func (fake *IpApiFake) Name() string {
	return "Fake IpAPi"
}

func (fake *IpApiFake) GetPublicIpAddress() (string, error) {
	return fake.Ip, fake.Error
}
