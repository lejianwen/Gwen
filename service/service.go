package service

type Service struct {
	AdminService
}

func New() *Service {
	all := new(Service)
	return all
}

var AllService = New()
