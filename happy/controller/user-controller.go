package controller

import (
	"happy/entity"
	"happy/service"

	"github.com/gin-gonic/gin"
)

type SaveUserController interface {
	UserFindAll() []entity.User
	UserSave(ctx *gin.Context) error
}

type userController struct {
	service service.UserService
}

func UserNew(service service.UserService) SaveUserController {
	return userController{
		service: service,
	}
}

func (c controller) UserFindAll() []entity.User {
	return c.service.FindAll()
}

func (c controller) UserSave(ctx *gin.Context) error {
	//TODO implement me
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	c.service.Save(user)
	return nil
}
