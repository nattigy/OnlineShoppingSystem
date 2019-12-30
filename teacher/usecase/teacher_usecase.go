package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/teacher"
)

type TeacherUsecase struct {
	teacherRepo teacher.TeacherRepository
}

func NewTeacherUsecase(TeacherRepo teacher.TeacherRepository) *TeacherUsecase {
	return &TeacherUsecase{teacherRepo: TeacherRepo}
}

func (t *TeacherUsecase) MakeNewPost(task models.Task, c models.ClassRoom) error {
	return nil
}

func (t *TeacherUsecase) EditPost(task models.Task) error {
	return nil
}

func (t *TeacherUsecase) RemoveTask(task models.Task) error {
	return nil
}

func (t *TeacherUsecase) UploadResource(resource models.Resource, s models.Subject, room models.ClassRoom) error {
	return nil
}

func (t *TeacherUsecase) TeacherUpdateProfile(teacher models.Teacher) error {
	return nil
}

func (t *TeacherUsecase) ReportGrade(grade models.Result, student models.Student) error {
	return nil
}

func (t *TeacherUsecase) ViewClasses(room models.ClassRoom) ([]models.Student, error) {
	return nil, nil
}
