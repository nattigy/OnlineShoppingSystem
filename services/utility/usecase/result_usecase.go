package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/utility"
)

type ResultUsecase struct {
	resultRepo utility.ResultRepository
}

func NewResultUsecase(resultRepo utility.ResultRepository) *ResultUsecase {
	return &ResultUsecase{resultRepo: resultRepo}
}

func (rr *ResultUsecase) Update(newResult models.Result) (models.Result, []error) {
	results, errs := rr.resultRepo.Update(newResult)
	return results, errs
}

func (rr *ResultUsecase) Delete(id uint) []error {
	errs := rr.resultRepo.Delete(id)
	return errs
}
