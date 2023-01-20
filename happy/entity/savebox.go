package entity

import (
	"time"
)

type SaveBox struct {
	ID            uint64    `gorm:"primaryKey" json:"box_id"`
	UserID        uint64    `json:"-"`
	BoxName       string    `json:"box-name" binding:"required" gorm:"type:varchar(100)"`
	CreateBoxDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create-date"`
	OpenBoxDate   time.Time `binding:"required" json:"open_date" gorm:"type:datetime"`
	Status        string    `json:"status" gorm:"type:varchar(100)"`
	IsOpened      bool      `json:"isopen" gorm:"derault:false"`
	SaveDiaries   []Diary   `json:"save-diaries" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:SaveBoxID;references:ID"`
}

// 다이어리 오픈 여부
func IsOpenedChange(openDate time.Time) bool {
	currentTime := time.Now()
	var isOpen bool

	// 오픈날짜가 아직 안왔으면
	if openDate.After(currentTime) {
		isOpen = false
	} else {
		isOpen = true
	}
	return isOpen
}
