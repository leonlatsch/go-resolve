package models

const (
	ProviderUpdateUrl    = "updateUrl"
	ProviderGoDaddy      = "goDaddy"
	ProviderHetzner      = "hetzner"
	ProviderHetznerCloud = "hetznerCloud"
)

type Config struct {
	Provider string   `json:"provider"`
	Interval string   `json:"interval"`
	Domain   string   `json:"domain"`
	Hosts    []string `json:"hosts"`
	OnlyUPNP bool     `json:"onlyUpnp"`

	UpdateUrlConfig    UpdateUrlConfig    `json:"updateUrlConfig"`
	GoDaddyConfig      GoDaddyConfig      `json:"goDaddyConfig"`
	HetznerConfig      HetznerConfig      `json:"hetznerConfig"`
	HetznerCloudConfig HetznerCloudConfig `json:"hetznerCloudConfig"`
}

type UpdateUrlConfig struct {
	Url string `json:"url"`
}

type GoDaddyConfig struct {
	ApiKey    string `json:"apiKey"`
	ApiSecret string `json:"apiSecret"`
}

type HetznerConfig struct {
	ZoneId   string `json:"zoneId"`
	ApiToken string `json:"apiToken"`
}

type HetznerCloudConfig struct {
	CloudApiToken string `json:"cloudApiToken"`
}

// Empty Config. Used for generating file at first launch

var EmptyConfig = Config{
	Provider: ProviderUpdateUrl,
	Interval: "1h",
	Domain:   "YOUR_DOMAIN",
	Hosts:    []string{"HOST1", "HOST2"},
	OnlyUPNP: false,

	UpdateUrlConfig: UpdateUrlConfig{
		Url: "UPDATE_URL",
	},
	GoDaddyConfig: GoDaddyConfig{
		ApiKey:    "GODADDY_API_KEY",
		ApiSecret: "GODADDY_API_SECRET",
	},
	HetznerConfig: HetznerConfig{
		ZoneId:   "ZONE_ID",
		ApiToken: "API_TOKEN",
	},
	HetznerCloudConfig: HetznerCloudConfig{
		CloudApiToken: "CLOUD_API_TOKEN",
	},
}
