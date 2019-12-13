package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/student"
)

type StudentUsecase struct {
	studentRepo student.StudentRepository
}

func NewStudentUsecase(StudentRepo student.StudentRepository) *StudentUsecase {
	return &StudentUsecase{studentRepo: StudentRepo}
}

func (st *StudentUsecase) ViewTasks(c models.ClassRoom, s models.Subject) ([]models.Task, error) {
	return nil, nil
}

func (st *StudentUsecase) Comment(t models.Task, student models.Student) error {
	return nil
}

func (st *StudentUsecase) StudentUpdateProfile(student models.Student) error {
	return nil
}

func (st *StudentUsecase) ViewClass(classRoom models.ClassRoom) ([]models.Student, error) {
	return nil, nil
}

func (st *StudentUsecase) ViewResources(subject models.Subject) ([]models.Resource, error) {
	return nil, nil
}

func (st *StudentUsecase) ViewResult(s models.Student) ([]models.Result, error) {
	return nil, nil
}
