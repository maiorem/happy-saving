package api

import (
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"happy/controller"
	"happy/middlewares"
	"happy/service"
	"io"
	"net/http"
	"os"
)

var (
	boxService    service.BoxService           = service.New()
	boxController controller.SaveBoxController = controller.New(boxService)
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

	server.GET("/boxlist", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, boxController.FindAll())
	})

	server.POST("/boxsave", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, boxController.Save(ctx))
	})

	server.Run(":8089")
}
