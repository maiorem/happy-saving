package dto

import "time"

type CreateUserRequest struct {
	BoxName  string
	Password string
	OpenDate time.Time
	Activate bool
	IsOpened bool
}
