package dto

import "time"

type CreateUserRequest struct {
	BoxName    string    `json:"box-name" binding:"required" gorm:"type:varchar(100)"`
	CreateDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create-date"`
	OpenDate   time.Time `binding:"required" gorm:"type:timestamp" json:"open-date"`
	Activate   bool      `json:"activate" gorm:"default:true"`
	IsOpened   bool      `json:"isopen" gorm:"derault:false"`
}
