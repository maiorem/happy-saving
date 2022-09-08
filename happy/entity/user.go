package entity

type User struct {
	Name  string    `json:"name"`
	Boxes []SaveBox `json:"boxes"`
}
