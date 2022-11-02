package dto

type CreateDiaryRequest struct {
	BoxID   uint64 `json:"box_id"`
	EmojiID uint64 `json:"emoji_id"`
	Content string `json:"content"`
}
