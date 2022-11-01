package service

import (
	"happy/repository"
)

type LoginService interface {
	Login(useremail string, password string) (uint64, bool)
}

type loginService struct {
	repository repository.Repository
}

func NewLoginService(happyRepository repository.Repository) LoginService {
	return &loginService{
		repository: happyRepository,
	}
}

func (service *loginService) Login(useremail string, password string) (uint64, bool) {
	var user = service.repository.Login(useremail)
	userid := user.ID
	return userid, password == user.Password
}
