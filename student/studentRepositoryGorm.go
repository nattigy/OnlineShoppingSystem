package student

import (
	"github.com/nattigy/parentschoolcommunicationsystem/gorm_models"
)

type StudentRepositoryGorm interface {
	ViewTasks(c gorm_models.ClassRoom, s gorm_models.Subject) ([]gorm_models.Task, error)
	Comment(t gorm_models.Task, student gorm_models.Student, d string) error
	StudentUpdateProfile(student gorm_models.Student) error
	ViewClass(classRoom gorm_models.ClassRoom) ([]gorm_models.Student, error)
	ViewResources(subject gorm_models.Subject) ([]gorm_models.Resources, error)
	ViewResult(s gorm_models.Student) ([]gorm_models.Result, error)
}
