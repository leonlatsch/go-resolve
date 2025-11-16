package service

type DnsModeService interface {
	UpdateDns(ip string) error
	Initialize() error
}
