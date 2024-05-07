package models

type Config struct {
	ApiKey    string   `json:"apiKey"`
	ApiSecret string   `json:"apiSecret"`
	Domain    string   `json:"domain"`
	Hosts     []string `json:"hosts"`
	Interval  string   `json:"interval"`
	UpdateUrl string   `json:"updateUrl"`
}
