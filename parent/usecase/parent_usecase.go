package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/parent"
)

type ParentUsecase struct {
	parentRepo parent.ParentRepository
}

func NewParentUsecase(parentRepo parent.ParentRepository) *ParentUsecase {
	return &ParentUsecase{parentRepo: parentRepo}
}

func (p *ParentUsecase) ViewGrade(student models.Student) (models.Result, error) {
	return models.Result{}, nil
}
