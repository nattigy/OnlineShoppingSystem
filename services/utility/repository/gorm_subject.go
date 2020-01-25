package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type SubjectRepository struct {
	conn *gorm.DB
}

func NewSubjectRepository(conn *gorm.DB) *SubjectRepository {
	return &SubjectRepository{conn: conn}
}

func (sr *SubjectRepository) CreateSubject(newSubject models.Subject) []error {
	errs := sr.conn.Create(&newSubject).GetErrors()
	return errs
}

func (sr *SubjectRepository) GetSubjectById(id uint) (models.Subject, []error) {
	subject := models.Subject{}
	errs := sr.conn.Where("id = ?", id).First(&subject).GetErrors()
	return subject, errs
}

func (sr *SubjectRepository) DeleteSubject(id uint) []error {
	errs := sr.conn.Unscoped().Where("id = ?", id).Delete(&models.Subject{}).GetErrors()
	return errs
}
