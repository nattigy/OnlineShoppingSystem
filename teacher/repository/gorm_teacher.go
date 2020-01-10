package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/gorm_models"
)

type GormTeacherRepository struct {
	conn *gorm.DB
}

func NewGormTeacherRepository(Conn *gorm.DB) *GormTeacherRepository {
	return &GormTeacherRepository{conn: Conn}
}

func (gs *GormTeacherRepository) MakeNewPost(task gorm_models.Task, c gorm_models.ClassRoom) error {
	task.ClassRoomId = c.Id
	gs.conn.Create(&task)
	return nil
}

func (gs *GormTeacherRepository) EditPost(task gorm_models.Task) error {
	gs.conn.Model(&task).Where("id = ?", 1).Updates(gorm_models.Task{Title: task.Title, Description: task.Description, ShortDescription: task.ShortDescription, Deadline: task.Deadline})
	return nil
}

func (gs *GormTeacherRepository) RemoveTask(task gorm_models.Task) error {
	gs.conn.Delete(&task)
	return nil
}

func (gs *GormTeacherRepository) UploadResource(resource gorm_models.Resources, s gorm_models.Subject, room gorm_models.ClassRoom) error {
	resources := gorm_models.Resources{SubjectId: s.Id, Title: resource.Title, Description: resource.Description, Path: resource.Path}
	gs.conn.Create(&resources)
	return nil
}

func (gs *GormTeacherRepository) TeacherUpdateProfile(teacher gorm_models.Teacher) error {
	gs.conn.Model(&teacher).Where("id = ?", 1).Updates(gorm_models.Teacher{FirstName: teacher.FirstName, MiddleName: teacher.MiddleName, Email: teacher.Email, Password: teacher.Password, ProfilePic: teacher.ProfilePic})
	return nil
}

func (gs *GormTeacherRepository) ReportGrade(grade gorm_models.Result, student gorm_models.Student) error {
	results := gorm_models.Result{StudentId: student.Id, Assessment: grade.Assessment, Final: grade.Final, Test: grade.Test, Total: grade.Total}
	gs.conn.Create(&results)
	return nil
}

func (gs *GormTeacherRepository) ViewClasses(room gorm_models.ClassRoom) ([]gorm_models.Student, error) {
	var class []gorm_models.Student
	gs.conn.Where("id = ?", 1).Find(&class)
	return class, nil
}
