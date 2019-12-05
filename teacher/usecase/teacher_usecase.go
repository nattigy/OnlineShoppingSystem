package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/teacher"
)

type TeacherUsecase struct {
	teacherRepo *teacher.TeacherRepository
}

func NewTeacherRepository(TeacherRepo *teacher.TeacherRepository) *TeacherUsecase {
	return &TeacherUsecase{teacherRepo: TeacherRepo}
}

func MakeNewPost(task models.Task) (bool, error) {
	return true, nil
}

func EditPost(task models.Task) (bool, error) {
	return false, nil
}

func RemoveTask(task models.Task) (bool, error) {
	return false, nil
}

func UploadResource(subject models.Subject) error {
	return nil
}

func TeacherUpdateProfile(teacher models.Teacher) error {
	return nil
}

func ReportGrade(grade models.Grade) error {
	return nil
}

func ViewClasses() ([]models.Student, error) {
	return nil, nil
}
