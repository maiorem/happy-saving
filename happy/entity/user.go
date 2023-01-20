package entity

import (
	"regexp"
	"time"
)

type User struct {
	ID         uint64    `gorm:"primaryKey" json:"user_id"`
	Name       string    `json:"name" binding:"required" gorm:"type:varchar(10)"`
	Email      string    `json:"email" validate:"required, email" gorm:"type:varchar(50)"`
	Password   string    `json:"password" validate:"required" gorm:"type:varchar(50)"`
	JoinDate   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"join-date" `
	IsActivate bool      `json:"is_activate" gorm:"type:bool; default:true"`
	Boxes      []SaveBox `json:"boxes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID;references:ID"`
}

// 이메일 유효성 검사 (bool)
func EmailValidationCheck(email string) bool {
	const emailPattern = `(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`
	matched, _ := regexp.MatchString(emailPattern, email)
	return matched
}

// 비밀번호 유효성 검사 (bool)
func PasswordValidationCheck(password string) bool {
	const pwPattern = `(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-])`
	matched, _ := regexp.MatchString(pwPattern, password)
	return matched
}

// 비밀번호 확인 (bool)
func PasswordConfirm(firstPw string, secondPw string) bool {
	return firstPw == secondPw
}
