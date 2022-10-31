package controller

import (
	"happy/dto"
	"happy/entity"
	"happy/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SaveBoxController interface {
	FindAll() []entity.SaveBox
	Save(ctx *gin.Context) error
	UpdateBox(ctx *gin.Context) error
	DeleteBox(ctx *gin.Context) error
	ActivateBox(ctx *gin.Context) entity.SaveBox
}

type boxController struct {
	service service.BoxService
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

func (c *boxController) ActivateBox(ctx *gin.Context) entity.SaveBox {
	return c.service.ActivateBox()

}

func (c *boxController) Save(ctx *gin.Context) error {
	//TODO implement me
	var savebox dto.CreateBoxRequest
	err := ctx.ShouldBindJSON(&savebox)
	if err != nil {
		return err
	}
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
