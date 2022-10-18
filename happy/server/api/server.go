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

	server.GET("/", func(ctx *gin.Context) {

	})

	server.GET("/boxlist", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, boxController.FindAll())
	})

	server.POST("/boxsave", func(ctx *gin.Context) {
		err := boxController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Box Input is Valid!!"})
		}

	})

	server.Run(":8089")
}
