package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type GormParentRepository struct {
	conn *gorm.DB
}

func NewGormParentRepository(Conn *gorm.DB) *GormParentRepository {
	return &GormParentRepository{conn: Conn}
}

func (g *GormParentRepository) ViewGrade(student models.Student) (models.Result, error) {
	return models.Result{}, nil
}
