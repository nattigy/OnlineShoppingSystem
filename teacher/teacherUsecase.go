package teacher

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type TeacherUsecase interface {
	MakeNewPost(task models.Task, classRoom models.ClassRoom) error
	EditPost(task models.Task) error
	RemoveTask(task models.Task) error
	UploadResource(resource models.Resources, s models.Subject, room models.ClassRoom) error
	TeacherUpdateProfile(teacher models.Teacher) error
	ReportGrade(grade models.Result, student models.Student) error
	ViewClasses(room models.ClassRoom) ([]models.Student, error)
}
