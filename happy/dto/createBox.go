package dto

import "time"

type CreateBoxRequest struct {
	UserID   uint64
	BoxName  string
	OpenDate time.Time
	Activate bool
	IsOpened bool
}
