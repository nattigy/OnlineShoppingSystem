package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type GormClassRoomRepository struct {
	conn *gorm.DB
}

func NewGormClassRoomRepository(conn *gorm.DB) *GormClassRoomRepository {
	return &GormClassRoomRepository{conn: conn}
}

func (cr *GormClassRoomRepository) CreateClassRoom(newClassRoom models.ClassRoom) []error {
	errs := cr.conn.Create(&newClassRoom).GetErrors()
	return errs
}

func (cr *GormClassRoomRepository) GetClassRoomById(id uint) (models.ClassRoom, []error) {
	classRoom := models.ClassRoom{}
	errs := cr.conn.Where("id = ?", id).Find(&classRoom).GetErrors()
	return classRoom, errs
}

func (cr *GormClassRoomRepository) DeleteClassRoom(id uint) []error {
	errs := cr.conn.Unscoped().Where("id = ?", id).Delete(&models.ClassRoom{}).GetErrors()
	return errs
}
