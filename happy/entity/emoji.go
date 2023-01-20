package entity

type Emoji struct {
	ID       uint64 `gorm:"primaryKey" json:"emoji_id"`
	Emoticon string `gorm:"type:longtext" json:"emoji"`
}
