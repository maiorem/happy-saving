package service

import (
	"happy-save-api/dto"
	"happy-save-api/entity"
	"happy-save-api/repository"
)

type DiaryService interface {
	FindAll() []entity.Diary
	Save(diary dto.CreateDiaryRequest) error
	FindById(id uint64) entity.Diary
	Update(diary dto.UpdateDiaryRequest) error
	DiaryCount(id uint64) int64

	EmojiAll() []entity.Emoji
	EmojiOne(id uint64) string
}

type diaryService struct {
	repository repository.Repository
}

func DiaryNew(diaryrepository repository.Repository) DiaryService {
	return &diaryService{
		repository: diaryrepository,
	}
}

func (d diaryService) FindAll() []entity.Diary {
	//TODO implement me
	return d.repository.FindAllDiary()
}

func (d diaryService) Save(diary dto.CreateDiaryRequest) error {
	//TODO implement me
	d.repository.Write(diary)
	return nil
}

func (d diaryService) FindById(id uint64) entity.Diary {
	//TODO implement me
	return d.repository.FindDiaryById(id)
}

func (d diaryService) Update(diary dto.UpdateDiaryRequest) error {
	//TODO implement me
	d.repository.UpdateDiary(diary)
	return nil
}

func (d diaryService) DiaryCount(id uint64) int64 {
	//TODO implement me
	return d.repository.CountDiaryInBox(id)
}

func (d diaryService) EmojiAll() []entity.Emoji {
	//TODO implement me
	return d.repository.FindAllEmoji()
}

func (d diaryService) EmojiOne(id uint64) string {
	return d.repository.EmojiOne(id)
}
