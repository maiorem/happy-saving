package controller

import (
	"happy-save-api/dto"
	"happy-save-api/entity"
	"happy-save-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SaveUserController interface {
	UserJoin(ctx *gin.Context) error
	UserFindAll() []entity.User
	FindByUserId(ctx *gin.Context) entity.User
	UserWithDraw(ctx *gin.Context) error
	UserUpdate(ctx *gin.Context) error
}

type userController struct {
	service    service.UserService
	jwtService service.JWTService
}

var userValidate *validator.Validate

func UserNew(service service.UserService) SaveUserController {
	validate = validator.New()
	return userController{
		service: service,
	}
}

func (c userController) UserFindAll() []entity.User {
	return c.service.FindAll()
}

func (c userController) FindByUserId(ctx *gin.Context) entity.User {
	//TODO implement me
	//id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	//if err != nil {
	//	ctx.JSON(http.StatusUnauthorized, "unauthorized")
	//}

	var metadata, err = service.NewJWTService().ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	userid, err := service.NewJWTService().FetchAuth(metadata)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.service.FindById(userid)
}

func (c userController) UserJoin(ctx *gin.Context) error {
	//TODO implement me
	var user dto.CreateUserRequest
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "invalid json")
		return err
	}

	c.service.Join(user)
	return nil
}

func (c userController) UserUpdate(ctx *gin.Context) error {
	//TODO implement me
	var updateuser dto.UpdateUserRequest
	err := ctx.ShouldBindJSON(&updateuser)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "invalid json")
		return err
	}

	metadata, err := service.NewJWTService().ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	userid, err := service.NewJWTService().FetchAuth(metadata)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	err = validate.Struct(updateuser)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "not validate")
		return err
	}

	updateuser.ID = userid

	c.service.Update(updateuser)
	return nil

}

func (c userController) UserWithDraw(ctx *gin.Context) error {
	//TODO implement me
	var user entity.User
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}

	user.ID = id
	err = validate.Struct(user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "not validate")
		return err
	}
	c.service.WithDrawUser(id)
	return nil
}
