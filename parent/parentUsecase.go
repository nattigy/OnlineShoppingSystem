package parent

import "github.com/nattigy/parentschoolcommunicationsystem/gorm_models"

type ParentUsecase interface {
	ViewGrade(student models.Student) ([]models.Result, error)
}
