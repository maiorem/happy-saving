package entity

import "time"

type Diary struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Box       SaveBox   `json:"box" binding:"required" gorm:"foreignkey:SaveBoxID"`
	SaveBoxID uint64    `json:"-"`
	Emoji     string    `json:"emoji" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"create-date" `
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"update-date"`
	Content   string    `json:"content" binding:"min=20, max=150" gorm:"type:varchar(150)"`
}
