package parent

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type ParentUsecase interface {
	ViewGrade(student models.Student) (models.Student, error)
}
