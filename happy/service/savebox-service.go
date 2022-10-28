package service

import (
	"happy/entity"
	"happy/repository"
)

type BoxService interface {
	Save(entity.SaveBox) error
	UpdateBox(entity.SaveBox) error
	DeleteBox(entity.SaveBox) error
	FindAll() []entity.SaveBox
	ActivateBox() entity.SaveBox
}

type boxService struct {
	repository repository.Repository
}

func BoxNew(boxRepository repository.Repository) BoxService {
	return &boxService{
		repository: boxRepository,
	}
}

func (service *boxService) Save(box entity.SaveBox) error {
	service.repository.Save(box)
	return nil
}

func (service *boxService) UpdateBox(box entity.SaveBox) error {
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

func (service *boxService) ActivateBox() entity.SaveBox {
	return service.repository.ActivateBox()
}
