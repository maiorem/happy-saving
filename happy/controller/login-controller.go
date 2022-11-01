package controller

import (
	"happy/dto"
	"happy/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func NewLoginController(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	isAuthenticated := controller.loginService.Login(credentials.Useremail, credentials.Password)
	if isAuthenticated {
		ts, err := controller.jWtService.GenerateToken(credentials.Useremail, true)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		ctx.JSON(http.StatusOK, tokens)

		saveErr := controller.jWtService.CreateAuth(credentials.Useremail, ts)
		if saveErr != nil {
			ctx.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		}
	}

}

// 로그아웃 로직 차후 검토 2022.11.01
func (controller *loginController) Logout(ctx *gin.Context) {
	metadata, err := controller.jWtService.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	delErr := controller.jWtService.DeleteTokens(metadata)
	if delErr != nil {
		ctx.JSON(http.StatusUnauthorized, delErr.Error())
		return
	}
	ctx.JSON(http.StatusOK, "Successfully logged out")
}
