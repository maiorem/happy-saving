package dto

type UpdateUserRequest struct {
	ID                 uint64 `json:"id"`
	Email              string `json:"email"`
	Name               string `json:"name"`
	OldPassowrd        string `json:"old_password"`
	NewPassword        string `json:"new_password"`
	NewConfirmPassword string `json:"new_password2"`
}
