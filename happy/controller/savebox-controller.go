package controller

import (
	"happy/dto"
	"happy/entity"
	"happy/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SaveBoxController interface {
	FindAll() []entity.SaveBox
	Save(ctx *gin.Context) error
	UpdateBox(ctx *gin.Context) error
	DeleteBox(ctx *gin.Context) error
	ActivateBox() entity.SaveBox
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

func (c *boxController) ActivateBox() entity.SaveBox {
	return c.service.ActivateBox()

}

func (c *boxController) Save(ctx *gin.Context) error {
	//TODO implement me
	var savebox dto.CreateBoxRequest
	err := ctx.ShouldBindJSON(&savebox)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "invalid json")
		return err
	}
	metadata, err := c.jwtService.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return err
	}

	userid, err := c.jwtService.FetchAuth(metadata)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return err
	}

	savebox.UserID = userid

	c.service.Save(savebox)
	return nil
}

func (c *boxController) UpdateBox(ctx *gin.Context) error {
	var box dto.UpdateBoxRequest
	err := ctx.ShouldBindJSON(&box)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	box.ID = id

	err = validate.Struct(box)
	if err != nil {
		return err
	}
	c.service.UpdateBox(box)
	return nil
}

func (c *boxController) DeleteBox(ctx *gin.Context) error {
	var box entity.SaveBox
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	box.ID = id
	err = validate.Struct(box)
	if err != nil {
		return err
	}

	c.service.DeleteBox(box)
	return nil
}
