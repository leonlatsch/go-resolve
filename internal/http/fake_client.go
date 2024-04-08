package http

type FakeHttpClient struct {
	RespBody string
	Error    error
}

func (self *FakeHttpClient) Get(url string, headers map[string]string) (string, error) {
	return self.RespBody, self.Error
}

func (self *FakeHttpClient) Put(url string, headers map[string]string, body interface{}) (string, error) {
	return self.RespBody, self.Error
}

func (self *FakeHttpClient) Patch(url string, headers map[string]string, body interface{}) (string, error) {
	return self.RespBody, self.Error
}
