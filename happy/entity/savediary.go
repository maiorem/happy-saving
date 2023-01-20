package entity

import (
	"time"
)

type Diary struct {
	ID          uint64    `gorm:"primaryKey" json:"diary_id"`
	SaveBoxID   uint64    `json:"-"`
	Emoji       Emoji     `json:"emoji" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:EmojiID;"`
	EmojiID     uint64    `json:"-"`
	CreatedDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create-date" `
	UpdatedDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update-date"`
	Content     string    `json:"content" binding:"min=20, max=150" gorm:"type:varchar(150)"`
}
