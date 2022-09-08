package service

import "happy/entity"

type BoxService interface {
	Save(box entity.SaveBox) entity.SaveBox
	FindAll() []entity.SaveBox
}

type boxService struct {
	boxes []entity.SaveBox
}

func New() BoxService {
	return &boxService{}
}

func (service *boxService) Save(box entity.SaveBox) entity.SaveBox {
	service.boxes = append(service.boxes, box)
	return box
}

func (service *boxService) FindAll() []entity.SaveBox {
	return service.boxes
}
