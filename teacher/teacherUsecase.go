package teacher

import "github.com/nattigy/parentschoolcommunicationsystem/gorm_models"

type TeacherUsecase interface {
	MakeNewPost(task models.Task) (bool, error)
	EditPost(task models.Task) (bool, error)
	RemoveTask(task models.Task) (bool, error)
	UploadResource(subject models.Subject) error
	TeacherUpdateProfile(teacher models.Teacher) error
	ReportGrade(grade models.Result) error
	ViewClasses() ([]models.Student, error)
}
