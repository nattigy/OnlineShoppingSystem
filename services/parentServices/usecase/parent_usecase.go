package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/parentServices"
)

type ParentUsecase struct {
	parentRepo parentServices.ParentRepository
}

func NewParentUsecase(parentRepo parentServices.ParentRepository) *ParentUsecase {
	return &ParentUsecase{parentRepo: parentRepo}
}

func (pr *ParentUsecase) AddParent(parent models.Parent) []error {
	errs := pr.parentRepo.AddParent(parent)
	return errs
}

func (pr *ParentUsecase) GetParents() ([]models.Parent, []error) {
	parents, errs := pr.parentRepo.GetParents()
	return parents, errs
}

func (pr *ParentUsecase) GetParentById(id uint) (models.Parent, []error) {
	parent, errs := pr.parentRepo.GetParentById(id)
	return parent, errs
}

func (pr *ParentUsecase) DeleteParent(id uint) []error {
	errs := pr.parentRepo.DeleteParent(id)
	return errs
}

func (pr *ParentUsecase) UpdateParent(newParent models.Parent) (models.Parent, []error) {
	parent, errs := pr.parentRepo.UpdateParent(newParent)
	return parent, errs
}

func (pr *ParentUsecase) ViewGrade(parentId uint) (models.Student, []error) {
	student, errs := pr.parentRepo.ViewGrade(parentId)
	return student, errs
}

func (pr *ParentUsecase) GetChild(parentId uint) (models.Student, []error) {
	student, errs := pr.GetChild(parentId)
	return student, errs
}
