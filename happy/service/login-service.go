package service

import (
	"happy/repository"
)

type LoginService interface {
	Login(useremail string, password string) bool
}

type loginService struct {
	repository repository.Repository
}

func NewLoginService(happyRepository repository.Repository) LoginService {
	return &loginService{
		repository: happyRepository,
	}
}

func (service *loginService) Login(useremail string, password string) bool {
	var user = service.repository.Login(useremail)
	return password == user.Password
}
