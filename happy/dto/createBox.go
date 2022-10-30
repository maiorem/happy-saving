package dto

import "time"

type CreateBoxRequest struct {
	BoxName  string
	OpenDate time.Time
	Activate bool
	IsOpened bool
}
