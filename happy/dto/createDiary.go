package dto

import (
	"happy/entity"
	"time"
)

type CreateDiaryRequest struct {
	Emoji     entity.Emoji
	UpdatedAt time.Time
	Content   string
}
