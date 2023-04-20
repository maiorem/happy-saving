package dto

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateNameRequest struct {
	Name string `json:"name"`
}
