package entity

import "time"

type Diary struct {
	Emoji      string    `json:"emoji"`
	CreateDate time.Time `json:"createDate"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
}
