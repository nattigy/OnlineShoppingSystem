package parent

import (
	"github.com/nattigy/parentschoolcommunicationsystem/gorm_models"
)

type ParentRepositoryGorm interface {
	ViewGrade(student gorm_models.Student) ([]gorm_models.Result, error)
}
