package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	parent2 "github.com/nattigy/parentschoolcommunicationsystem/services/parent"
)

type ParentUsecase struct {
	parentRepo parent2.ParentRepository
}

func NewParentUsecase(parentRepo parent2.ParentRepository) *ParentUsecase {
	return &ParentUsecase{parentRepo: parentRepo}
}

func (p *ParentUsecase) ViewGrade(student models.Student) ([]models.Result, error) {
	data, err := p.parentRepo.ViewGrade(student)
	if err != nil {
		return data, err
	}
	return data, nil
}
