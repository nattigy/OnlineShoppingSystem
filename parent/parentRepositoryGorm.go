package parent

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type ParentRepositoryGorm interface {
	ViewGrade(student models.Student) ([]models.Result, error)
}
