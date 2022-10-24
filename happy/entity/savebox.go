package entity

import "time"

type SaveBox struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Author      User      `json:"author" binding:"required" gorm:"foreignkey:UserID"`
	UserID      uint64    `json:"-"`
	BoxName     string    `json:"box-name" binding:"required" gorm:"type:varchar(100)"`
	CreateDate  time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"create-date"`
	OpenDate    time.Time `json:"-" binding:"required" gorm:"type:timestamp" json:"open-date"`
	SaveDiaries []Diary   `json:"save-diaries"`
}
