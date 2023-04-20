package api

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"happy-save-api/controller"
	"happy-save-api/middlewares"
	"happy-save-api/repository"
	"happy-save-api/service"
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

	userService    service.UserService           = service.UserNew(happyRepository)
	userController controller.SaveUserController = controller.UserNew(userService)

	diaryService    service.DiaryService           = service.DiaryNew(happyRepository)
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

	server.GET("/", happymain) // main test

	// 이메일 인증
	server.POST("/verify-email", verifyEmail)

	// health check
	server.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Health check!"})
	})

	server.POST("/login", userlogin)
	server.POST("/join", userjoin)

	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := server.Group("/api")
	{

		apiRoutes.POST("/refresh", refresh) // 토큰 리프레시

		apiRoutes.POST("/logout", userlogout)    // 로그아웃
		apiRoutes.POST("/name", createName)      // 이름 생성
		apiRoutes.GET("/user", userOne)          // 내 정보 보기
		apiRoutes.PUT("/user", userModify)       // 이름, 비밀번호 수정
		apiRoutes.PUT("/withdraw", userWithDraw) // 회원 탈퇴 (회원 계정 비활성화)

		apiRoutes.POST("/box", myboxsave)         // 저금통 생성
		apiRoutes.GET("/user/box", myboxlist)     // 저금통 목록
		apiRoutes.PUT("/box/:id", myboxupdate)    // 저금통 수정 (이름)
		apiRoutes.DELETE("/box/:id", myboxdelete) // 저금통 삭제 (완전삭제?)
		apiRoutes.GET("/box/:id", myboxone)       // history 저금통 상세
		apiRoutes.GET("/box/now", nowbox)         // 현재 작성 중 저금통 (메인)

		//apiRoutes.GET("/box/:id", mydiarylist)        // 저금통 상세 (다이어리 목록)
		apiRoutes.POST("/diary", writediary)          // 다이어리 작성
		apiRoutes.GET("/diary/:id", mydiaryById)      // 다이어리 상세
		apiRoutes.PUT("/diary/:id", mydiaryupdate)    // 다이어리 내용 수정
		apiRoutes.GET("/diary/:id/count", countdiary) // 현재 저금통 내 다이어리 수

		apiRoutes.GET("/emoji", emojilist) // 이모지 리스트
		apiRoutes.GET("/emoji/:id", emoji) // 이모지 단건

	}

	port := "8000"
	server.Run(":" + port)

}

// 이메일 인증
func verifyEmail(c *gin.Context) {
	// 인증 링크를 보낼 사용자 이메일 주소
	toEmail := c.Query("email")

	// gomail 패키지를 사용하여 이메일 전송 설정
	mail := gomail.NewMessage()
	mail.SetHeader("To", toEmail)
	mail.SetHeader("Subject", "임시 비밀번호가 발급되었습니다.")
	mail.SetBody("text/html", "<a href='http://localhost:8080/verified?email="+toEmail+"'>인증하기</a>")

	// SMTP 서버 설정
	d := gomail.NewDialer("smtp.gmail.com", 587, "maiorem00", "Mireena0510!")

	// 이메일 전송
	if err := d.DialAndSend(mail); err != nil {
		fmt.Println("Error sending email:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "임시 비밀번호가 발급되었습니다. 이메일을 확인해주세요."})
}

func refresh(ctx *gin.Context) {
	loginController.Refresh(ctx)
}

// 메인
func happymain(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Health check OK!"})
}

// ================================= 회원
func userlogin(ctx *gin.Context) {
	loginController.Login(ctx)
}

func userlogout(ctx *gin.Context) {
	loginController.Logout(ctx)
}

func userjoin(ctx *gin.Context) {

	err := userController.UserJoin(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "New Member join!"})
	}

}

func createName(ctx *gin.Context) {

	err := userController.UserCreateName(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "New Name Create!"})
	}

}

func userOne(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, userController.FindByUserId(ctx))
}

func userModify(ctx *gin.Context) {
	err := userController.UserUpdate(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "User Update"})
	}

}

func userWithDraw(ctx *gin.Context) {
	err := userController.UserWithDraw(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "User Delete"})
	}

}

// ================================= 저금통
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

func myboxupdate(context *gin.Context) {
	err := boxController.UpdateBox(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"message": "Box update! "})
	}

}

func myboxdelete(context *gin.Context) {
	err := boxController.DeleteBox(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"message": "Box Delete! "})
	}

}

func nowbox(context *gin.Context) {
	context.JSON(http.StatusOK, boxController.FindNotOpenedBox(context))
}

func myboxone(context *gin.Context) {
	context.JSON(http.StatusOK, boxController.FindByIdBox(context))

}

// ================================= 일기
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
	ctx.JSON(http.StatusOK, diaryController.DiaryFindById(ctx))

}

func mydiaryupdate(ctx *gin.Context) {
	err := diaryController.DiaryUpdate(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Diary update! "})
	}

}

func countdiary(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, diaryController.DiaryCount(ctx))
}

func emojilist(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, diaryController.EmojiAll())
}

func emoji(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, diaryController.EmojiOne(ctx))
}
