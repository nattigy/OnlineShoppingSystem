package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices"
)

type TeacherUsecase struct {
	teacherRepo teacherServices.TeacherRepository
}

func NewTeacherUsecase(teacherRepo teacherServices.TeacherRepository) *TeacherUsecase {
	return &TeacherUsecase{teacherRepo: teacherRepo}
}

func (tr *TeacherUsecase) AddTeacher(newTeacher models.Teacher) []error {
	errs := tr.teacherRepo.AddTeacher(newTeacher)
	return errs
}

func (tr *TeacherUsecase) GetTeachers() ([]models.Teacher, []error) {
	teachers, errs := tr.teacherRepo.GetTeachers()
	return teachers, errs
}

func (tr *TeacherUsecase) GetTeacherById(id uint) (models.Teacher, []error) {
	teacher, errs := tr.teacherRepo.GetTeacherById(id)
	return teacher, errs
}

func (tr *TeacherUsecase) DeleteTeacher(id uint) []error {
	errs := tr.teacherRepo.DeleteTeacher(id)
	return errs
}

func (tr *TeacherUsecase) UpdateTeacher(newTeacher models.Teacher) (models.Teacher, []error) {
	teacher, errs := tr.teacherRepo.UpdateTeacher(newTeacher)
	return teacher, errs
}

func (tr *TeacherUsecase) CreateTask(task models.Task) []error {
	errs := tr.teacherRepo.CreateTask(task)
	return errs
}

func (tr *TeacherUsecase) GetTasks(subjectId uint) ([]models.Task, []error) {
	tasks, errs := tr.teacherRepo.GetTasks(subjectId)
	return tasks, errs
}

func (tr *TeacherUsecase) UpdateTask(newTask models.Task) (models.Task, []error) {
	task, errs := tr.teacherRepo.UpdateTask(newTask)
	return task, errs
}

func (tr *TeacherUsecase) DeleteTask(taskId uint) []error {
	errs := tr.teacherRepo.DeleteTask(taskId)
	return errs
}

func (tr *TeacherUsecase) UploadResource(resource models.Resources) []error {
	errs := tr.teacherRepo.UploadResource(resource)
	return errs
}

func (tr *TeacherUsecase) DeleteResource(resourceId uint) []error {
	errs := tr.teacherRepo.DeleteResource(resourceId)
	return errs
}

func (tr *TeacherUsecase) ReportGrade(grade models.Result) []error {
	errs := tr.teacherRepo.ReportGrade(grade)
	return errs
}

func (tr *TeacherUsecase) ViewStudents(classRoomId uint) ([]models.Student, []error) {
	students, errs := tr.teacherRepo.ViewStudents(classRoomId)
	return students, errs
}

func (tr *TeacherUsecase) GetResource(subjectId uint) ([]models.Resources, []error) {
	respurces, errs := tr.teacherRepo.GetResource(subjectId)
	return respurces, errs
}
