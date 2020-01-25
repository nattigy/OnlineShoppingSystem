package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type ResultRepository struct {
	conn *gorm.DB
}

func NewResultRepository(conn *gorm.DB) *ResultRepository {
	return &ResultRepository{conn: conn}
}

func (rr *ResultRepository) Update(newResult models.Result) (models.Result, []error) {
	result := models.Result{}
	errs := rr.conn.Model(&result).Where("id = ?", newResult.Id).Updates(&newResult).GetErrors()
	return result, errs
}

func (rr *ResultRepository) Delete(id uint) []error {
	errs := rr.conn.Unscoped().Where("id = ?", id).Delete(&models.Result{}).GetErrors()
	return errs
}
