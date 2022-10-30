package dto

import (
	"happy/entity"
	"time"
)

type UpdateDiaryRequest struct {
	Emoji     entity.Emoji
	UpdatedAt time.Time
	Content   string
}
