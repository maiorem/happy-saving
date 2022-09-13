package entity

type User struct {
	Name  string    `json:"name" binding:"required"`
	Email string    `json:"email" validate:"required, email"`
	Boxes []SaveBox `json:"boxes"`
}
