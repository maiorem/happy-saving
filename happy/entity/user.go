package entity

import "time"

type User struct {
	ID       uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UUID     string    `json:"uuid" binding:"required" gorm:"type:varchar(100)"`
	Name     string    `json:"name" binding:"required" gorm:"type:varchar(10)"`
	Email    string    `json:"email" validate:"required, email" gorm:"type:varchar(50)"`
	JoinDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"join-date" `
	Boxes    []SaveBox `json:"boxes"`
}
