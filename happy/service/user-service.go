package service

import "happy/entity"

type UserService interface {
	// FindById(uuid string) entity.User
	FindAll() []entity.User
	Save(user entity.User) entity.User
	// Update(user entity.User) entity.User
}

type userService struct {
	users []entity.User
}

func UserNew() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.User) entity.User {
	service.users = append(service.users, user)
	return user
}

func (service *userService) FindAll() []entity.User {
	return service.users
}
