package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/utility"
)

type ClassRoomUsecase struct {
	classRoomRepo utility.ClassRoomRepository
}

func NewClassRoomUsecase(classRoomRepo utility.ClassRoomRepository) *ClassRoomUsecase {
	return &ClassRoomUsecase{classRoomRepo: classRoomRepo}
}

func (cr *ClassRoomUsecase) CreateClassRoom(newClassRoom models.ClassRoom) []error {
	errs := cr.classRoomRepo.CreateClassRoom(newClassRoom)
	return errs
}

func (cr *ClassRoomUsecase) GetClassRoomById(id uint) (models.ClassRoom, []error) {
	classRoom, errs := cr.classRoomRepo.GetClassRoomById(id)
	return classRoom, errs
}

func (cr *ClassRoomUsecase) DeleteClassRoom(id uint) []error {
	errs := cr.classRoomRepo.DeleteClassRoom(id)
	return errs
}
