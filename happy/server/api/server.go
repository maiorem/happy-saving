package api

import (
	"happy/controller"
	"happy/middlewares"
	"happy/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	boxService      service.BoxService             = service.BoxNew()
	boxController   controller.SaveBoxController   = controller.BoxNew(boxService)
	userService     service.UserService            = service.UserNew()
	userController  controller.SaveUserController  = controller.UserNew(userService)
	diaryService    service.DiaryService           = service.DiaryNew()
	diaryController controller.SaveDiaryController = controller.DiaryNew(diaryService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func Start() {

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	server.GET("/", happymain)

	server.GET("/boxlist", myboxlist)

	server.POST("/boxsave", myboxsave)

	server.POST("/join", userjoin)

	server.POST("/writediary", writediary)

	server.GET("/diaries", mydiarylist)

	server.GET("/diaries/:id", mydiaryById)

	server.Run(":8089")
}

func happymain(ctx *gin.Context) {
	// 회원 정보가 없으면

	// 회원 정보가 존재하면
	ctx.JSON(http.StatusOK, boxController.FindAll())
}

func myboxlist(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, boxController.FindAll())
}

func myboxsave(ctx *gin.Context) {
	err := boxController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Box Input is Valid!!"})
	}
}

func userjoin(ctx *gin.Context) {

	err := userController.UserSave(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "New Member join!"})
	}

}

func writediary(ctx *gin.Context) {
	err := diaryController.DiarySave(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Write Diary!"})
	}

}

func mydiarylist(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, diaryController.DiaryFindAll())
}

func mydiaryById(ctx *gin.Context) {
	// id := ctx.Param("id")

}
