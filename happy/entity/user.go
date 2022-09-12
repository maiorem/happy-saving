package entity

type User struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Boxes []SaveBox `json:"boxes"`
}
