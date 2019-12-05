package student

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type StudentUsecase interface {
	ViewTasks(subject models.Student) ([]models.Task, error)
	Comment(t models.Task, student models.Student) ([]models.Comment, error)
	StudentUpdateProfile(student models.Student) error
	ViewClass(student models.Student) ([]models.Student, error)
	ViewResources(subject models.Subject) ([]models.Resource, error)
}
