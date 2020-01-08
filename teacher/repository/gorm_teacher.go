package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type GormTeacherRepository struct {
	conn *gorm.DB
}

func NewGormTeacherRepository(Conn *gorm.DB) *GormTeacherRepository {
	return &GormTeacherRepository{conn: Conn}
}

func (gs *GormTeacherRepository) MakeNewPost(task models.Task, c models.ClassRoom) error {
	return nil
}

func (gs *GormTeacherRepository) EditPost(task models.Task) error {
	return nil
}

func (gs *GormTeacherRepository) RemoveTask(task models.Task) error {
	return nil
}

func (gs *GormTeacherRepository) UploadResource(resource models.Resource, s models.Subject, room models.ClassRoom) error {
	return nil
}

func (gs *GormTeacherRepository) TeacherUpdateProfile(teacher models.Teacher) error {
	return nil
}

func (gs *GormTeacherRepository) ReportGrade(grade models.Result, student models.Student) error {
	return nil
}

func (gs *GormTeacherRepository) ViewClasses(room models.ClassRoom) ([]models.Student, error) {
	return nil, nil
}
