package entity

type User struct {
	UUID string `json:"uuid" binding:"required"`
	Name  string    `json:"name" binding:"required"`
	Email string    `json:"email" validate:"required, email"`
	Boxes []SaveBox `json:"boxes"`
}
