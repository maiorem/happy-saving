package controller

import (
	"github.com/gin-gonic/gin"
	"happy/entity"
	"happy/service"
)

type SaveBoxController interface {
	FindAll() []entity.SaveBox
	Save(ctx *gin.Context) entity.SaveBox
}

type controller struct {
	service service.BoxService
}

func New(service service.BoxService) SaveBoxController {
	return controller{
		service: service,
	}
}
func (c controller) FindAll() []entity.SaveBox {
	//TODO implement me
	return c.service.FindAll()
}

func (c controller) Save(ctx *gin.Context) entity.SaveBox {
	//TODO implement me
	var savebox entity.SaveBox
	ctx.ShouldBindJSON(&savebox)
	c.service.Save(savebox)
	return savebox
}
