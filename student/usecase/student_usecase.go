package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/gorm_models"
	"github.com/nattigy/parentschoolcommunicationsystem/student"
)

type StudentUsecase struct {
	studentRepo student.StudentRepository
}

func NewStudentUsecase(StudentRepo student.StudentRepository) *StudentUsecase {
	return &StudentUsecase{studentRepo: StudentRepo}
}

func (st *StudentUsecase) ViewTasks(c models.ClassRoom, s models.Subject) ([]models.Task, error) {
	data, err := st.studentRepo.ViewTasks(c, s)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (st *StudentUsecase) Comment(t models.Task, student models.Student, data string) error {
	err := st.studentRepo.Comment(t, student, data)
	if err != nil {
		return err
	}
	return nil
}

func (st *StudentUsecase) StudentUpdateProfile(student models.Student) error {
	err := st.studentRepo.StudentUpdateProfile(student)
	if err != nil {
		return err
	}
	return nil
}

func (st *StudentUsecase) ViewClass(classRoom models.ClassRoom) ([]models.Student, error) {
	data, err := st.studentRepo.ViewClass(classRoom)
	if err != nil {
		return data, err
	}
	return nil, nil
}

func (st *StudentUsecase) ViewResources(subject models.Subject) ([]models.Resource, error) {
	data, err := st.studentRepo.ViewResources(subject)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (st *StudentUsecase) ViewResult(s models.Student) ([]models.Result, error) {
	data, err := st.studentRepo.ViewResult(s)
	if err != nil {
		return data, err
	}
	return data, nil
}
