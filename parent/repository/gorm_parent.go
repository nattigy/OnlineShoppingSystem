package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/gorm_models"
)

type GormParentRepository struct {
	conn *gorm.DB
}

func NewGormParentRepository(Conn *gorm.DB) *GormParentRepository {
	return &GormParentRepository{conn: Conn}
}

func (g *GormParentRepository) ViewGrade(student gorm_models.Student) ([]gorm_models.Result, error) {
	var result []gorm_models.Result
	g.conn.Where("student_id = ?", 1).Find(&result)
	return result, nil
}
