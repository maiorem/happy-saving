package service

import "happy/entity"

type DiaryService interface {
	FindAll() []entity.Diary
	Save(diary entity.Diary) entity.Diary
	// FindById(diary_id string) entity.Diary
	// Update(diary entity.Diary)
}

type diaryService struct {
	diaries []entity.Diary
}

func DiaryNew() DiaryService {
	return &diaryService{}
}

func (service *diaryService) Save(diary entity.Diary) entity.Diary {
	service.diaries = append(service.diaries, diary)
	return diary
}

func (service *diaryService) FindAll() []entity.Diary {
	return service.diaries
}
