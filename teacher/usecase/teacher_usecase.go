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

func (t *TeacherUsecase) MakeNewPost(task models.Task, subject models.Subject) error {
	err := t.MakeNewPost(task, subject)
	return err

}

func (t *TeacherUsecase) EditPost(task models.Task) error {
	err := t.EditPost(task)
	return err
}

func (t *TeacherUsecase) RemoveTask(task models.Task) error {
	err := t.RemoveTask(task)
	return err
}

func (t *TeacherUsecase) UploadResource(resource models.Resources, s models.Subject, room models.ClassRoom) error {
	err := t.UploadResource(resource, s, room)
	return err
}

func (t *TeacherUsecase) TeacherUpdateProfile(teacher models.Teacher) error {
	err := t.TeacherUpdateProfile(teacher)
	return err
}

func (t *TeacherUsecase) ReportGrade(grade models.Result, student models.Student) error {
	err := t.ReportGrade(grade, student)
	return err
}

func (t *TeacherUsecase) ViewClasses(room models.ClassRoom) ([]models.Student, error) {
	data, err := t.ViewClasses(room)
	return data, err
}
