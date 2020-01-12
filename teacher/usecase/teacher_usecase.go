package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/teacher"
)

type TeacherUsecase struct {
	teacherRepo teacher.TeacherRepository
}

func NewTeacherUsecase(teacherRepo teacher.TeacherRepository) *TeacherUsecase {
	return &TeacherUsecase{teacherRepo: teacherRepo}
}

func (t *TeacherUsecase) MakeNewPost(task models.Task, subject models.Subject) []error {
	err := t.teacherRepo.MakeNewPost(task, subject)
	return err
}

func (t *TeacherUsecase) EditPost(task models.Task) []error {
	err := t.teacherRepo.EditPost(task)
	return err
}

func (t *TeacherUsecase) RemoveTask(task models.Task) []error {
	err := t.teacherRepo.RemoveTask(task)
	return err
}

func (t *TeacherUsecase) UploadResource(resource models.Resources, s models.Subject, room models.ClassRoom) []error {
	err := t.teacherRepo.UploadResource(resource, s, room)
	return err
}

func (t *TeacherUsecase) TeacherUpdateProfile(teacher models.Teacher) []error {
	err := t.teacherRepo.TeacherUpdateProfile(teacher)
	return err
}

func (t *TeacherUsecase) ReportGrade(grade models.Result, student models.Student) []error {
	err := t.teacherRepo.ReportGrade(grade, student)
	return err
}

func (t *TeacherUsecase) ViewClasses(room models.ClassRoom) ([]models.Student, []error) {
	data, err := t.teacherRepo.ViewClasses(room)
	return data, err
}

func (t *TeacherUsecase) FetchPosts(subject models.Subject) ([]models.Task, []error) {
	tasks, errs := t.teacherRepo.FetchPosts(subject)
	return tasks, errs
}
