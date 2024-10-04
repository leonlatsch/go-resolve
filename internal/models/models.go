package models

const (
	ProviderUpdateUrl = "updateUrl"
	ProviderGoDaddy   = "goDaddy"
)

type Config struct {
	Provider string   `json:"provider"`
	Interval string   `json:"interval"`
	Domain   string   `json:"domain"`
	Hosts    []string `json:"hosts"`

	UpdateUrlConfig UpdateUrlConfig `json:"updateUrlConfig"`
	GoDaddyConfig   GoDaddyConfig   `json:"goDaddyConfig"`
}

type UpdateUrlConfig struct {
	Url string `json:"url"`
}

type GoDaddyConfig struct {
	ApiKey    string `json:"apiKey"`
	ApiSecret string `json:"apiSecret"`
}

// Empty Config. Used for generating file at first launch

var EmptyConfig = Config{
	Provider: ProviderUpdateUrl,
	Interval: "1h",
	Domain:   "YOUR_DOMAIN",
	Hosts:    []string{"HOST1", "HOST2"},

	UpdateUrlConfig: UpdateUrlConfig{
		Url: "UPDATE_URL",
	},
	GoDaddyConfig: GoDaddyConfig{
		ApiKey:    "GODADDY_API_KEY",
		ApiSecret: "GODADDY_API_SECRET",
	},
}
