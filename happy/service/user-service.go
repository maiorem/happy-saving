package service

import (
	"happy-save-api/dto"
	"happy-save-api/entity"
	"happy-save-api/repository"
)

type UserService interface {
	Join(user dto.CreateUserRequest) error
	Update(user dto.UpdateUserRequest) error
	Login(email string) error
	FindById(id uint64) entity.User
	FindAll() []entity.User
	WithDrawUser(userid uint64) error
	CreateName(userid uint64, name dto.CreateNameRequest) error
}

type userService struct {
	repository repository.Repository
}

func UserNew(userRepository repository.Repository) UserService {
	return &userService{
		repository: userRepository,
	}
}

func (service *userService) Join(user dto.CreateUserRequest) error {
	service.repository.Join(user)
	return nil
}

func EmailValidationCheck(email string) error {

	return nil
}

func PasswordValidationCheck(password string) error {

	return nil
}
func (service *userService) CreateName(userid uint64, name dto.CreateNameRequest) error {
	//TODO implement me
	service.repository.CreateName(userid, name)
	return nil
}
func PasswordConfirm(firstPw string, secondPw string) error {

	return nil
}

func (service *userService) Login(email string) error {
	//TODO implement me
	service.repository.Login(email)
	return nil
}

func (service *userService) FindById(id uint64) entity.User {
	//TODO implement me
	return service.repository.FindById(id)
}

func (service *userService) Update(user dto.UpdateUserRequest) error {
	//TODO implement me
	service.repository.UpdateUser(user)
	return nil
}

func (service *userService) FindAll() []entity.User {
	return service.repository.FindAllUser()
}

func (service *userService) WithDrawUser(userid uint64) error {
	//TODO implement me
	service.repository.DeleteUser(userid)
	return nil
}
