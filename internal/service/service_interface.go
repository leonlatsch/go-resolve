package service

type DnsModeService interface {
	ObserveAndUpdateDns()
	UpdateDns(ip string)
	Initialize() error
}
