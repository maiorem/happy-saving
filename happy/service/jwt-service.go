package service

import (
	"errors"
	"fmt"
	"log"

	"net/http"
	"strconv"
	"strings"
	"time"

	"happy-save-api/common"
	"happy-save-api/dto"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type JWTService interface {
	GenerateToken(userid uint64, admin bool) (td dto.TokenDetails, err error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	VerifyToken(r *http.Request) (*jwt.Token, error)
	CreateAuth(userId uint64, td dto.TokenDetails) (err error)
	DeleteTokens(authD *dto.AccessDetails) error
	ExtractTokenMetadata(r *http.Request) (*dto.AccessDetails, error)
	FetchAuth(authD *dto.AccessDetails) (uint64, error)
	DeleteAuth(givenUuid string) (uint64, error)
}

type jwtACCESSCustomClaims struct {
	UserID     uint64 `json:"user_id"`
	Admin      bool   `json:"admin"`
	AccessUUID string `json:"access_uuid"`
	jwt.StandardClaims
}

type jwtREFRESHCustomClaims struct {
	UserID      uint64 `json:"user_id"`
	Admin       bool   `json:"admin"`
	RefreshUUID string `json:"refresh_uuid"`
	jwt.StandardClaims
}

type jwtService struct {
	ACCESS_SECRET  string
	REFRESH_SECRET string
	issuer         string
}

func NewJWTService() JWTService {
	return &jwtService{
		ACCESS_SECRET:  GetAccessSecret(),
		REFRESH_SECRET: GetRefreshSecret(),
		issuer:         "www.koldsleep.com",
	}
}

func GetAccessSecret() string {
	return "AccessEnfant21878!"
}

func GetRefreshSecret() string {
	return "RefreshEnfatn21878!"
}

// 토큰 생성
func (jwtSrv *jwtService) GenerateToken(userid uint64, admin bool) (td dto.TokenDetails, err error) {

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = fmt.Sprintf("%s++%d", td.AccessUuid, userid)

	// Set custom and standard claims
	atClaims := &jwtACCESSCustomClaims{
		userid,
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
	td.AccessToken, err = at.SignedString([]byte(jwtSrv.ACCESS_SECRET))
	if err != nil {
		panic(err)
	}

	rtClaims := &jwtREFRESHCustomClaims{
		userid,
		admin,
		td.RefreshUuid,
		jwt.StandardClaims{
			ExpiresAt: td.RtExpires,
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(jwtSrv.REFRESH_SECRET))
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
		return []byte(jwtSrv.ACCESS_SECRET), nil
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
		return []byte(jwtSrv.ACCESS_SECRET), nil
	})
}

// 메타데이터 출력 차후 검토 2022.11.01
func (jwtSrv *jwtService) ExtractTokenMetadata(r *http.Request) (*dto.AccessDetails, error) {
	token, err := jwtSrv.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if !ok {
			return nil, err
		}
		return &dto.AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

// 레디스에 토큰 저장
func (jwtSrv *jwtService) CreateAuth(userId uint64, td dto.TokenDetails) (err error) {
	client := common.GetClient()

	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	if err = client.Set(td.AccessUuid, strconv.Itoa(int(userId)), at.Sub(now)).Err(); err != nil {
		return
	}
	if err = client.Set(td.RefreshUuid, strconv.Itoa(int(userId)), rt.Sub(now)).Err(); err != nil {
		return
	}

	return
}

// 레디스 토큰 삭제 로직 차후 검토  2022.11.01
func (jwtSrv *jwtService) DeleteTokens(authD *dto.AccessDetails) error {
	client := common.GetClient()

	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%d", authD.AccessUuid, authD.UserId)

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

func (jwtSrv *jwtService) FetchAuth(authD *dto.AccessDetails) (uint64, error) {
	client := common.GetClient()
	
	log.Println("auth ID: ", authD.AccessUuid)

	userid, err := client.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}

	userID, _ := strconv.ParseUint(userid, 10, 64)
	if authD.UserId != userID {
		return 0, errors.New("unauthorized")
	}
	return userID, nil
}

func (jwtSrv *jwtService) DeleteAuth(givenUuid string) (uint64, error) {
	client := common.GetClient()
	deleted, err := client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}

	return uint64(deleted), nil
}
