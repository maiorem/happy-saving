package entity

type SaveBox struct {
	BoxName    string `json:"box-name"`
	CreateDate string `json:"create-date"`
	OpenDate   string `json:"open-date"`

	SaveDiaries []Diary `json:"save-diaries"`
}
