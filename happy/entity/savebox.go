package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type SaveBox struct {
	gorm.Model
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID      uint64    `json:"-"`
	BoxName     string    `json:"box-name" binding:"required" gorm:"type:varchar(100)"`
	CreateDate  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create-date"`
	OpenDate    time.Time `binding:"required" gorm:"type:timestamp" json:"open-date"`
	Activate    bool      `json:"activate" gorm:"default:true"`
	IsOpened    bool      `json:"isopen" gorm:"derault:false"`
	SaveDiaries []Diary   `json:"save-diaries" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignkey:SaveBoxID"`
}

// 다이어리 오픈 여부
func IsOpenedChange(openDate time.Time, isOpen bool) bool {
	currentTime := time.Now()

	// 오픈날짜가 아직 안왔으면
	if openDate.After(currentTime) {
		isOpen = false
	} else {
		isOpen = true
	}
	return isOpen
}
