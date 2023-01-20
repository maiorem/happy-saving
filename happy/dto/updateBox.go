package dto

type UpdateBoxRequest struct {
	ID       uint64 `json:"box_id"`
	UserID   uint64 `json:"user_id"`
	BoxName  string `json:"box_name"`
	OpenDate string `json:"open_date"`
}
