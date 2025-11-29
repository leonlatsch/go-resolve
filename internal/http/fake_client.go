package http

type FakeHttpClient struct {
	RespBody string
	Error    error
}

func (fake *FakeHttpClient) Get(url string, headers map[string]string) (string, error) {
	return fake.RespBody, fake.Error
}

func (fake *FakeHttpClient) Put(url string, headers map[string]string, body any) (string, error) {
	return fake.RespBody, fake.Error
}

func (fake *FakeHttpClient) Post(url string, headers map[string]string, body any) (string, error) {
	return fake.RespBody, fake.Error
}

func (fake *FakeHttpClient) Patch(url string, headers map[string]string, body any) (string, error) {
	return fake.RespBody, fake.Error
}
