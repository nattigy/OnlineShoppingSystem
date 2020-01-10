package student

import "github.com/nattigy/parentschoolcommunicationsystem/gorm_models"

type StudentRepositoryGorm interface {
	ViewTasks(c models.ClassRoom, s models.Subject) ([]models.Task, error)
	Comment(t models.Task, student models.Student, d string) error
	StudentUpdateProfile(student models.Student) error
	ViewClass(classRoom models.ClassRoom) ([]models.Student, error)
	ViewResources(subject models.Subject) ([]models.Resource, error)
	ViewResult(s models.Student) ([]models.Result, error)
}
