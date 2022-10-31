package service

type LoginService interface {
	Login(useremail string, password string) bool
}

type loginService struct {
	authorizedUserEmail string
	authorizedPassword  string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUserEmail: "maiorem",
		authorizedPassword:  "123456",
	}
}

func (service *loginService) Login(useremail string, password string) bool {
	return service.authorizedUserEmail == useremail &&
		service.authorizedPassword == password
}
