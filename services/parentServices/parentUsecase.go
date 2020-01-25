package parentServices

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type ParentUsecase interface {
	AddParent(parent models.Parent) []error
	GetParents() ([]models.Parent, []error)
	GetParentById(id uint) (models.Parent, []error)
	DeleteParent(id uint) []error
	UpdateParent(newParent models.Parent) (models.Parent, []error)
	ViewGrade(parentId uint) (models.Student, []error)
	GetChild(parentId uint) (models.Student, []error)
}
