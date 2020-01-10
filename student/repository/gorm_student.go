package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type GormStudentRepository struct {
	conn *gorm.DB
}

func NewGormStudentRepository(Conn *gorm.DB) *GormStudentRepository {
	return &GormStudentRepository{conn: Conn}
}

func (gs *GormStudentRepository) ViewTasks(classRoom models.ClassRoom, subject models.Subject) ([]models.Task, error) {
	var task []models.Task
	gs.conn.Where("class_room_id = ? AND subject_id = ?", 1, 1).Find(&task)
	return task, nil
}

func (gs *GormStudentRepository) Comment(task models.Task, student models.Student, d string) error {
	comment := models.Comment{TaskId: task.Id, StudentId: student.Id, Data: d}
	gs.conn.Create(&comment)
	return nil
}

func (gs *GormStudentRepository) StudentUpdateProfile(student models.Student) error {
	gs.conn.Model(&student).Where("id = ?", 1).Updates(models.Student{Email: student.Email, Password: student.Password, ProfilePic: student.ProfilePic})
	return nil
}

func (gs *GormStudentRepository) ViewClass(classRoom models.ClassRoom) ([]models.Student, error) {
	var student []models.Student
	gs.conn.Where("class_room_id = ?", 1).Find(&student)
	return student, nil
}

func (gs *GormStudentRepository) ViewResources(subject models.Subject) ([]models.Resources, error) {
	var resources []models.Resources
	gs.conn.Where("subject_id = ?", 1).Find(&resources)
	return resources, nil
}

func (gs *GormStudentRepository) ViewResult(student models.Student) ([]models.Result, error) {
	var result []models.Result
	gs.conn.Where("student_id = ?", 1).Find(&result)
	return result, nil
}
