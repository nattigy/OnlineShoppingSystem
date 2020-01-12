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

func (gs *GormStudentRepository) ViewTasks(classRoom models.ClassRoom, subject models.Subject) ([]models.Task, []error) {
	var task []models.Task
	errs := gs.conn.Where("class_room_id = ? AND subject_id = ?", classRoom.Id, subject.Id).Find(&task).GetErrors()
	return task, errs
}

func (gs *GormStudentRepository) Comment(task models.Task, student models.Student, d string) []error {
	comment := models.Comment{TaskId: task.Id, StudentId: student.Id, Data: d}
	errs := gs.conn.Create(&comment).GetErrors()
	return errs
}

func (gs *GormStudentRepository) StudentUpdateProfile(student models.Student) []error {
	errs := gs.conn.Model(&student).Where("id = ?", student.Id).Updates(models.Student{Email: student.Email, Password: student.Password, ProfilePic: student.ProfilePic}).GetErrors()
	return errs
}

func (gs *GormStudentRepository) ViewClass(classRoom models.ClassRoom) ([]models.Student, []error) {
	var student []models.Student
	errs := gs.conn.Where("class_room_id = ?", classRoom.Id).Find(&student).GetErrors()
	return student, errs
}

func (gs *GormStudentRepository) ViewResources(subject models.Subject) ([]models.Resources, []error) {
	var resources []models.Resources
	errs := gs.conn.Where("subject_id = ?", subject.Id).Find(&resources).GetErrors()
	return resources, errs
}

func (gs *GormStudentRepository) ViewResult(student models.Student) ([]models.Result, []error) {
	var result []models.Result
	errs := gs.conn.Where("student_id = ?", student.Id).Find(&result).GetErrors()
	return result, errs
}
