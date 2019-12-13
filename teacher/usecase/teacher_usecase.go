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

func MakeNewPost(task models.Task, c models.ClassRoom) error {
	return nil
}

func EditPost(task models.Task) error {
	return nil
}

func RemoveTask(task models.Task) error {
	return nil
}

func UploadResource(resource models.Resource, s models.Subject, room models.ClassRoom) error {
	return nil
}

func TeacherUpdateProfile(teacher models.Teacher) error {
	return nil
}

func ReportGrade(grade models.Result, student models.Student) error {
	return nil
}

func ViewClasses(room models.ClassRoom) ([]models.Student, error) {
	return nil, nil
}
