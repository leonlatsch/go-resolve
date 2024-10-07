package hetzner

type ZoneId string
type RecordId string

type Record struct {
	Id     RecordId `json:"id"`
	Value  string   `json:"value"`
	Type   string   `json:"type"`
	Name   string   `json:"name"`
	ZoneId ZoneId   `json:"zone_id"`
}

type RecordsWrapper struct {
	Records []Record `json:"records"`
}
