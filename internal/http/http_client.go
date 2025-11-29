package http

type HttpClient interface {
	Get(url string, headers map[string]string) (string, error)
	Put(url string, headers map[string]string, body any) (string, error)
	Post(url string, headers map[string]string, body any) (string, error)
	Patch(url string, headers map[string]string, body any) (string, error)
}
