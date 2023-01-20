package service

import (
	"happy-save-api/dto"
	"happy-save-api/entity"
	"happy-save-api/repository"
)

type BoxService interface {
	Save(dto.CreateBoxRequest) error
	UpdateBox(dto.UpdateBoxRequest) error
	DeleteBox(entity.SaveBox) error
	FindAll() []entity.SaveBox
	ActivateBox(uint64) entity.SaveBox
	FindById(id uint64) entity.SaveBox
}

type boxService struct {
	repository repository.Repository
}

func BoxNew(boxRepository repository.Repository) BoxService {
	return &boxService{
		repository: boxRepository,
	}
}

func (service *boxService) Save(box dto.CreateBoxRequest) error {
	service.repository.Save(box)
	return nil
}

func (service *boxService) UpdateBox(box dto.UpdateBoxRequest) error {
	service.repository.UpdateBox(box)
	return nil
}

func (service *boxService) DeleteBox(box entity.SaveBox) error {
	service.repository.DeleteBox(box)
	return nil
}

func (service *boxService) FindAll() []entity.SaveBox {
	return service.repository.FindAllBox()
}

// 메인페이지 (박스별 일기장 리스트)
func (service *boxService) ActivateBox(userId uint64) entity.SaveBox {
	var box = service.repository.ActivateBox(userId)
	box.IsOpened = entity.IsOpenedChange(box.OpenBoxDate)
	box.SaveDiaries = service.repository.FindAllDiaryByBoxId(box.ID)
	return box
}

func (service *boxService) FindById(id uint64) entity.SaveBox {
	//TODO implement me
	return service.repository.FindHistoryBoxById(id)
}
