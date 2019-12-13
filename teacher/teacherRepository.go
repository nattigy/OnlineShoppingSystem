package teacher

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type TeacherRepository interface {
	MakeNewPost(task models.Task, c models.ClassRoom) error
	EditPost(task models.Task) error
	RemoveTask(task models.Task) error
	UploadResource(resource models.Resource, s models.Subject, room models.ClassRoom) error
	TeacherUpdateProfile(teacher models.Teacher) error
	ReportGrade(grade models.Result, student models.Student) error
	ViewClasses(room models.ClassRoom) ([]models.Student, error)
}
