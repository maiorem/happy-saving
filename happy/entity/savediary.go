package entity

type Diary struct {
	Emoji      string `json:"emoji"`
	CreateDate string `json:"createDate"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}
