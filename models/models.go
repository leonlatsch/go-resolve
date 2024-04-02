package models

type DnsRecord struct {
	Data string `json:"data"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Config struct {
	ApiKey    string   `json:"apiKey"`
	ApiSecret string   `json:"apiSecret"`
	Domain    string   `json:"domain"`
	Hosts     []string `json:"hosts"`
	Interval  string   `json:"interval"`
}
