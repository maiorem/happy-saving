package dto

type UpdateUserRequest struct {
	ID       uint64
	Name     string
	Password string
	Email    string
}
