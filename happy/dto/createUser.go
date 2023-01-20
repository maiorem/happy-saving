package dto

type CreateUserRequest struct {
	Email           string `json:"email" validate:"required"`
	Name            string `json:"name" binding:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"password2" validate:"required"`
}
