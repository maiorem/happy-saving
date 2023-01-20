package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"happy-save-api/dto"
	"happy-save-api/entity"
	"happy-save-api/service"
	"net/http"
)

type SaveBoxController interface {
	FindAll() []entity.SaveBox
	Save(ctx *gin.Context) error
	UpdateBox(ctx *gin.Context) error
	DeleteBox(ctx *gin.Context) error
	FindByIdBox(ctx *gin.Context) entity.SaveBox
	ActivateBox(ctx *gin.Context) entity.SaveBox
	FindNotOpenedBox(context *gin.Context) entity.SaveBox
}

type boxController struct {
	service    service.BoxService
	jwtService service.JWTService
}

var validate *validator.Validate

func BoxNew(service service.BoxService) SaveBoxController {
	validate = validator.New()
	return &boxController{
		service: service,
	}
}
func (c *boxController) FindAll() []entity.SaveBox {
	//TODO implement me
	return c.service.FindAll()
}
func (c *boxController) FindNotOpenedBox(ctx *gin.Context) entity.SaveBox {

	metadata, err := service.NewJWTService().ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	userid, err := service.NewJWTService().FetchAuth(metadata)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.service.ActivateBox(userid)
}

func (c *boxController) ActivateBox(ctx *gin.Context) entity.SaveBox {
	metadata, err := service.NewJWTService().ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	userid, err := service.NewJWTService().FetchAuth(metadata)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.service.ActivateBox(userid)

}

func (c *boxController) Save(ctx *gin.Context) error {
	//TODO implement me
	var savebox dto.CreateBoxRequest
	err := ctx.ShouldBindJSON(&savebox)
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

	savebox.UserID = userid

	err = c.service.Save(savebox)
	if err != nil {
		return err
	}
	return nil
}

func (c *boxController) UpdateBox(ctx *gin.Context) error {
	var box dto.UpdateBoxRequest
	err := ctx.ShouldBindJSON(&box)
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

	box.UserID = userid

	err = validate.Struct(box)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "not validate")
		return err
	}
	err = c.service.UpdateBox(box)
	if err != nil {
		return err
	}
	return nil
}

func (c *boxController) DeleteBox(ctx *gin.Context) error {
	var box entity.SaveBox
	metadata, err := service.NewJWTService().ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	userid, err := service.NewJWTService().FetchAuth(metadata)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	box.UserID = userid
	err = validate.Struct(box)
	if err != nil {
		return err
	}

	err = c.service.DeleteBox(box)
	if err != nil {
		return err
	}
	return nil
}

func (c *boxController) FindByIdBox(ctx *gin.Context) entity.SaveBox {
	//TODO implement me
	metadata, err := service.NewJWTService().ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	userid, err := service.NewJWTService().FetchAuth(metadata)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.service.FindById(userid)

}
