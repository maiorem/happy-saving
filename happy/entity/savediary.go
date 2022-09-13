package entity

import "time"

type Diary struct {
	Emoji      string    `json:"emoji"`
	CreateDate time.Time `json:"createDate"`
	Content    string    `json:"content" binding:"min=20, max=50"`
}
