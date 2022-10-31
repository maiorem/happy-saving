package dto

import "time"

type UpdateBoxRequest struct {
	ID       uint64
	BoxName  string
	OpenDate time.Time
	Activate bool
	IsOpened bool
}
