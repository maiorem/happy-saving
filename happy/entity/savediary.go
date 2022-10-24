package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Diary struct {
	gorm.Model
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	SaveBoxID uint64    `json:"-"`
	Emoji     Emoji     `json:"emoji" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignkey:EmojiID"`
	EmojiID   uint64    `json:"-"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create-date" `
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update-date"`
	Content   string    `json:"content" binding:"min=20, max=150" gorm:"type:varchar(150)"`
}
type Emoji struct {
	gorm.Model
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Emoticon string `gorm:"type:varchar(20)" json:"emoji"`
}
