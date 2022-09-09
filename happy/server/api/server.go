package api

import (
	"github.com/gin-gonic/gin"
	"happy/controller"
	"happy/middlewares"
	"happy/service"
	"net/http"
)

var (
	boxService    service.BoxService           = service.New()
	boxController controller.SaveBoxController = controller.New(boxService)
)

func Start() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/boxlist", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, boxController.FindAll())
	})

	server.POST("/boxsave", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, boxController.Save(ctx))
	})

	server.Run(":8089")
}
