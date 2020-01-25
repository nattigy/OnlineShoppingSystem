package teacherServices

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type TeacherUsecase interface {
	AddTeacher(newTeacher models.Teacher) []error
	GetTeachers() ([]models.Teacher, []error)
	GetTeacherById(id uint) (models.Teacher, []error)
	DeleteTeacher(id uint) []error
	UpdateTeacher(newTeacher models.Teacher) (models.Teacher, []error)
	CreateTask(task models.Task) []error
	GetTasks(subjectId uint) ([]models.Task, []error)
	UpdateTask(newTask models.Task) (models.Task, []error)
	DeleteTask(taskId uint) []error
	UploadResource(resource models.Resources) []error
	DeleteResource(resourceId uint) []error
	ReportGrade(grade models.Result) []error
	ViewStudents(classRoomId uint) ([]models.Student, []error)
}
