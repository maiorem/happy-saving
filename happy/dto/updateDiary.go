package dto

import (
	"happy/entity"
	"time"
)

type UpdateDiaryRequest struct {
	ID        uint64
	Emoji     entity.Emoji
	UpdatedAt time.Time
	Content   string
}
