package entity

import (
	"regexp"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID       uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UUID     string    `json:"uuid" binding:"required" gorm:"type:varchar(100)"`
	Name     string    `json:"name" binding:"required" gorm:"type:varchar(10)"`
	Email    string    `json:"email" validate:"required, email" gorm:"type:varchar(50)"`
	Password string    `json:"password" validate:"required" gorm:"type:varchar(50)"`
	JoinDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"join-date" `
	Boxes    []SaveBox `json:"boxes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignkey:UserID"`
}

// 이메일 유효성 검사 (bool)
func EmailValidationCheck(email string) bool {
	const emailPattern = `(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`
	matched, _ := regexp.MatchString(emailPattern, email)
	return matched
}

// 비밀번호 유효성 검사 (bool)

// 비밀번호 확인 (bool)

// 비밀번호 변경 (return newPassword)
