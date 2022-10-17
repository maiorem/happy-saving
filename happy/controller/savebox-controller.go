package controller

import (
	"happy/entity"
	"happy/service"

	"github.com/gin-gonic/gin"
)

type SaveBoxController interface {
	FindAll() []entity.SaveBox
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.BoxService
}

func BoxNew(service service.BoxService) SaveBoxController {
	return controller{
		service: service,
	}
}
func (c controller) FindAll() []entity.SaveBox {
	//TODO implement me
	return c.service.FindAll()
}

func (c controller) Save(ctx *gin.Context) error {
	//TODO implement me
	var savebox entity.SaveBox
	err := ctx.ShouldBindJSON(&savebox)
	if err != nil {
		return err
	}
	c.service.Save(savebox)
	return nil
}
