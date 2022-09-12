package entity

import "time"

type SaveBox struct {
	BoxName     string    `json:"box-name" binding:"required"`
	CreateDate  time.Time `json:"create-date"`
	OpenDate    time.Time `json:"open-date"`
	SaveDiaries []Diary   `json:"save-diaries"`
}
