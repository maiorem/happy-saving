package dto

import "time"

type UpdateBoxRequest struct {
	BoxName  string
	OpenDate time.Time
	Activate bool
	IsOpened bool
}
