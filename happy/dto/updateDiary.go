package dto

type UpdateDiaryRequest struct {
	ID      uint64 `json:"diary_id"`
	BoxID   uint64 `json:"box_id"`
	EmojiID uint64 `json:"emoji_id"`
	Content string `json:"content"`
}
