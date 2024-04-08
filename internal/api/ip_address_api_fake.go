package api

type IpApiFake struct {
	Ip    string
	Error error
}

func (self *IpApiFake) GetPublicIpAddress() (string, error) {
	return self.Ip, self.Error
}
