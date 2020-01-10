package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/gorm_models"
)

type GormStudentRepository struct {
	conn *gorm.DB
}

func NewGormStudentRepository(Conn *gorm.DB) *GormStudentRepository {
	return &GormStudentRepository{conn: Conn}
}

func (gs *GormStudentRepository) ViewTasks(c models.ClassRoom, s models.Subject) ([]models.Task, error) {
	return nil, nil
}

func (gs *GormStudentRepository) Comment(t models.Task, s models.Student, d string) error {
	return nil
}

func (gs *GormStudentRepository) StudentUpdateProfile(student models.Student) error {
	return nil
}

func (gs *GormStudentRepository) ViewClass(classRoom models.ClassRoom) ([]models.Student, error) {
	return nil, nil
}

func (gs *GormStudentRepository) ViewResources(subject models.Subject) ([]models.Resource, error) {
	return nil, nil
}

func (gs *GormStudentRepository) ViewResult(s models.Student) ([]models.Result, error) {
	return nil, nil
}
