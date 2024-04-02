package models

type DnsRecord struct {
	Data string `json:"data"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type DomainDetail struct {
	Domain       string        `json:"domain"`
	ContactAdmin DomainContact `json:"contactAdmin"`
}

type DomainContact struct {
	Email     string `json:"email"`
	FirstName string `json:"nameFirst"`
	LastName  string `json:"nameLast"`
}
