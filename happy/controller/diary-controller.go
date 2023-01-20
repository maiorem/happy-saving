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

type SaveDiaryController interface {
	DiaryFindAll() []entity.Diary
	DiarySave(ctx *gin.Context) error
	DiaryFindById(ctx *gin.Context) entity.Diary
	DiaryUpdate(ctx *gin.Context) error
	DiaryCount(ctx *gin.Context) int64

	EmojiAll() []entity.Emoji
	EmojiOne(ctx *gin.Context) string
}

type diaryController struct {
	service    service.DiaryService
	jwtService service.JWTService
}

var diaryvalidate *validator.Validate

func DiaryNew(service service.DiaryService) SaveDiaryController {
	validate = validator.New()
	return diaryController{
		service: service,
	}
}

func (c diaryController) DiaryFindAll() []entity.Diary {
	return c.service.FindAll()
}

func (c diaryController) DiarySave(ctx *gin.Context) error {
	var diary dto.CreateDiaryRequest
	err := ctx.ShouldBindJSON(&diary)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "invalid json")
		return err
	}

	c.service.Save(diary)
	return nil
}

func (c diaryController) DiaryFindById(ctx *gin.Context) entity.Diary {
	//TODO implement me
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}
	return c.service.FindById(id)
}

func (c diaryController) DiaryUpdate(ctx *gin.Context) error {
	//TODO implement me
	var diaryupdate dto.UpdateDiaryRequest
	err := ctx.ShouldBindJSON(&diaryupdate)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	diaryupdate.ID = id

	err = validate.Struct(diaryupdate)
	if err != nil {
		return err
	}
	c.service.Update(diaryupdate)
	return nil
}

func (c diaryController) DiaryCount(ctx *gin.Context) int64 {
	//TODO implement me
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}
	return c.service.DiaryCount(id)
}

func (c diaryController) EmojiAll() []entity.Emoji {
	//TODO implement me
	return c.service.EmojiAll()
}

func (c diaryController) EmojiOne(ctx *gin.Context) string {
	//TODO implement me
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}
	return c.service.EmojiOne(id)
}
