package parent

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type ParentRepository interface {
	ViewGrade(student models.Student) (models.Student, error)
}
