package dto

import "time"

type CreateBoxRequest struct {
	UserID   uint64    `json:"user_id"`
	BoxName  string    `json:"box_name"`
	OpenDate time.Time `json:"open_date"`
	Activate bool      `json:"activate"`
	IsOpened bool      `json:"is_open"`
}
