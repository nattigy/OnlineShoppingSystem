package teacher

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type TeacherRepository interface {
	MakeNewPost(task models.Task) (bool, error)
	EditPost(task models.Task) (bool, error)
	RemoveTask(task models.Task) (bool, error)
	UploadResource(subject models.Subject) error
	TeacherUpdateProfile(teacher models.Teacher) error
	ReportGrade(grade models.Grade) error
	ViewClasses() ([]models.Student, error)
}
