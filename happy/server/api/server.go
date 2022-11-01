package api

import (
	"happy/controller"
	"happy/middlewares"
	"happy/repository"
	"happy/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	happyRepository repository.Repository = repository.NewRepository()

	loginService    service.LoginService       = service.NewLoginService(happyRepository)
	jwtService      service.JWTService         = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)

	boxService    service.BoxService           = service.BoxNew(happyRepository)
	boxController controller.SaveBoxController = controller.BoxNew(boxService)

	userService    service.UserService           = service.UserNew()
	userController controller.SaveUserController = controller.UserNew(userService)

	diaryService    service.DiaryService           = service.DiaryNew()
	diaryController controller.SaveDiaryController = controller.DiaryNew(diaryService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func Start() {

	defer happyRepository.CloseDB()

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.CORSmiddleware())

	server.GET("/", happymain)

	server.POST("/login", userlogin)
	server.POST("/join", userjoin)

	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{

		apiRoutes.GET("/boxlist", myboxlist)
		apiRoutes.POST("/boxsave", myboxsave)

		apiRoutes.POST("/writediary", writediary)

		apiRoutes.GET("/diaries", mydiarylist)

		apiRoutes.GET("/diaries/:id", mydiaryById)
	}

	// We can setup this env variable from the EB console
	port := os.Getenv("PORT")

	// Elastic Beanstalk forwards requests to port 8089
	if port == "" {
		port = "8089"
	}
	server.Run(":" + port)

}

func happymain(ctx *gin.Context) {
	// 회원 정보가 없으면

	// 회원 정보가 존재하면
	ctx.JSON(http.StatusOK, boxController.ActivateBox())
}

func userlogin(ctx *gin.Context) {
	loginController.Login(ctx)
}

func userjoin(ctx *gin.Context) {

	err := userController.UserSave(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "New Member join!"})
	}

}

func myboxlist(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, boxController.FindAll())
}

func myboxsave(ctx *gin.Context) {

	err := boxController.Save(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "New Box is Valid"})
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
