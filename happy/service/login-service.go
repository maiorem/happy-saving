package service

import (
	"happy-save-api/repository"
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

func (service *loginService) Login(email string, password string) (uint64, bool) {
	var user = service.repository.Login(email)
	userid := user.ID
	return userid, password == user.Password
}
