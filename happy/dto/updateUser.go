package dto

import "time"

type UpdateUserRequest struct {
	BoxName  string
	Password string
	OpenDate time.Time
	Activate bool
	IsOpened bool
}
