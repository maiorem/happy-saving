package controller

import (
	"fmt"
	"happy-save-api/dto"
	"happy-save-api/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Refresh(c *gin.Context)
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

	userId, isAuthenticated := controller.loginService.Login(credentials.Useremail, credentials.Password)
	if isAuthenticated {
		ts, err := controller.jWtService.GenerateToken(userId, true)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}

		saveErr := controller.jWtService.CreateAuth(userId, ts)
		if saveErr != nil {
			ctx.JSON(http.StatusUnprocessableEntity, saveErr.Error())
			return
		}

		ctx.JSON(http.StatusOK, tokens)
		return

	} else {
		ctx.JSON(http.StatusUnauthorized, "비밀번호 틀림")
		return
	}

}

// 로그아웃 로직 차후 검토 2022.11.01
func (controller *loginController) Logout(ctx *gin.Context) {
	metadata, err := controller.jWtService.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized / "+err.Error())
		return
	}

	delErr := controller.jWtService.DeleteTokens(metadata)
	if delErr != nil {
		ctx.JSON(http.StatusUnauthorized, delErr.Error())
		return
	}
	ctx.JSON(http.StatusOK, "Successfully logged out")
}

func (controller *loginController) Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.GetRefreshSecret()), nil
	})

	//if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}

	//is token valid?
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "user_id Error occurred / "+err.Error())
			return
		}
		admin, err := strconv.ParseBool(fmt.Sprintf("%t", claims["admin"]))
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "admin Error occurred / "+err.Error())
			return
		}
		//Delete the previous Refresh Token
		deleted, delErr := service.NewJWTService().DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			c.JSON(http.StatusUnauthorized, "unauthorized ")
			return
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := service.NewJWTService().GenerateToken(userId, admin)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}
		//save the tokens metadata to redis
		saveErr := service.NewJWTService().CreateAuth(userId, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)

		return
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired / "+err.Error())
		return
	}
}
