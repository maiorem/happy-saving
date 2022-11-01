package service

import (
	"errors"
	"fmt"

	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"happy/common"
	"happy/dto"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type JWTService interface {
	GenerateToken(name string, admin bool) (td dto.TokenDetails, err error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	VerifyToken(r *http.Request) (*jwt.Token, error)
	CreateAuth(useremail string, td dto.TokenDetails) (err error)
	DeleteTokens(authD *dto.AccessDetails) error
}

type jwtCustomClaims struct {
	Useremail string `json:"email"`
	Admin     bool   `json:"admin"`
	UUID      string `json:"token_uuid"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "www.koldsleep.com",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (jwtSrv *jwtService) GenerateToken(useremail string, admin bool) (td dto.TokenDetails, err error) {

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = fmt.Sprintf("%s++%s", td.AccessUuid, useremail)

	// Set custom and standard claims
	atClaims := &jwtCustomClaims{
		useremail,
		admin,
		td.AccessUuid,
		jwt.StandardClaims{
			ExpiresAt: td.AtExpires,
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create token with claims
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	// Generate encoded token using the secret signing key
	td.AccessToken, err = at.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}

	rtClaims := &jwtCustomClaims{
		useremail,
		admin,
		td.RefreshUuid,
		jwt.StandardClaims{
			ExpiresAt: td.RtExpires,
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}

	return td, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(jwtSrv.secretKey), nil
	})
}

func (jwtSrv *jwtService) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(jwtSrv.secretKey), nil
	})
}

// 레디스에 토큰 저장
func (jwtSrv *jwtService) CreateAuth(useremail string, td dto.TokenDetails) (err error) {
	client := common.GetClient()

	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	if err = client.Set(td.AccessUuid, strconv.Quote(useremail), at.Sub(now)).Err(); err != nil {
		return
	}
	if err = client.Set(td.RefreshUuid, strconv.Quote(useremail), rt.Sub(now)).Err(); err != nil {
		return
	}

	return
}

// 레디스 토큰 삭제 로직 차후 검토  2022.11.01
func (jwtSrv *jwtService) DeleteTokens(authD *dto.AccessDetails) error {
	client := common.GetClient()

	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%s", authD.AccessUuid, authD.UserEmail)

	//delete access token
	deletedAt, err := client.Del(authD.AccessUuid).Result()
	if err != nil {
		return err
	}

	//delete refresh token
	deletedRt, err := client.Del(refreshUuid).Result()
	if err != nil {
		return err
	}

	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}

	return nil
}
