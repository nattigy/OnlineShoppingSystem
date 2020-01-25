package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/utility"
)

type SubjectUsecase struct {
	subjectRepo utility.SubjectRepository
}

func NewSubjectUsecase(subjectRepo utility.SubjectRepository) *SubjectUsecase {
	return &SubjectUsecase{subjectRepo: subjectRepo}
}

func (sr *SubjectUsecase) CreateSubject(newSubject models.Subject) []error {
	errs := sr.subjectRepo.CreateSubject(newSubject)
	return errs
}

func (sr *SubjectUsecase) GetSubjectById(id uint) (models.Subject, []error) {
	subject, errs := sr.subjectRepo.GetSubjectById(id)
	return subject, errs
}

func (sr *SubjectUsecase) DeleteSubject(id uint) []error {
	errs := sr.subjectRepo.DeleteSubject(id)
	return errs
}
