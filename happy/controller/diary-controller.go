package controller

import (
	"happy/entity"
	"happy/service"

	"github.com/gin-gonic/gin"
)

type SaveDiaryController interface {
	DiaryFindAll() []entity.Diary
	DiarySave(ctx *gin.Context) error
}

type diaryController struct {
	service service.DiaryService
}

func DiaryNew(service service.DiaryService) SaveDiaryController {
	return diaryController{
		service: service,
	}
}

func (c diaryController) DiaryFindAll() []entity.Diary {
	return c.service.FindAll()
}

func (c diaryController) DiarySave(ctx *gin.Context) error {
	var diary entity.Diary
	err := ctx.ShouldBindJSON(&diary)
	if err != nil {
		return err
	}
	c.service.Save(diary)
	return nil
}
