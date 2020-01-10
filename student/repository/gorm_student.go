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

func (gs *GormStudentRepository) ViewTasks(classRoom gorm_models.ClassRoom, subject gorm_models.Subject) ([]gorm_models.Task, error) {
	var task []gorm_models.Task
	gs.conn.Where("class_room_id = ? AND subject_id = ?", 1, 1).Find(&task)
	return task, nil
}

func (gs *GormStudentRepository) Comment(task gorm_models.Task, student gorm_models.Student, d string) error {
	comment := gorm_models.Comment{TaskId: task.ID, StudentId: student.Id, Data: d}
	gs.conn.Create(&comment)
	return nil
}

func (gs *GormStudentRepository) StudentUpdateProfile(student gorm_models.Student) error {
	gs.conn.Model(&student).Where("id = ?", 1).Updates(gorm_models.Student{Email: student.Email, Password: student.Password, ProfilePic: student.ProfilePic})
	return nil
}

func (gs *GormStudentRepository) ViewClass(classRoom gorm_models.ClassRoom) ([]gorm_models.Student, error) {
	var student []gorm_models.Student
	gs.conn.Where("class_room_id = ?", 1).Find(&student)
	return student, nil
}

func (gs *GormStudentRepository) ViewResources(subject gorm_models.Subject) ([]gorm_models.Resources, error) {
	var resources []gorm_models.Resources
	gs.conn.Where("subject_id = ?", 1).Find(&resources)
	return resources, nil
}

func (gs *GormStudentRepository) ViewResult(student gorm_models.Student) ([]gorm_models.Result, error) {
	var result []gorm_models.Result
	gs.conn.Where("student_id = ?", 1).Find(&result)
	return result, nil
}
