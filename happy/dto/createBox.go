package dto

type CreateBoxRequest struct {
	UserID   uint64 `json:"user_id" binding:"required"`
	BoxName  string `json:"box_name" binding:"required"`
	OpenDate string `json:"open_date" binding:"required"`
}
