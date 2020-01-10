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
	task.ClassRoomId = c.Id
	gs.conn.Create(&task)
	return nil
}

func (gs *GormTeacherRepository) EditPost(task models.Task) error {
	gs.conn.Model(&task).Where("id = ?", 1).Updates(models.Task{Title: task.Title, Description: task.Description, ShortDescription: task.ShortDescription, Deadline: task.Deadline})
	return nil
}

func (gs *GormTeacherRepository) RemoveTask(task models.Task) error {
	gs.conn.Delete(&task)
	return nil
}

func (gs *GormTeacherRepository) UploadResource(resource models.Resources, s models.Subject, room models.ClassRoom) error {
	resources := models.Resources{SubjectId: s.Id, Title: resource.Title, Description: resource.Description, Path: resource.Path}
	gs.conn.Create(&resources)
	return nil
}

func (gs *GormTeacherRepository) TeacherUpdateProfile(teacher models.Teacher) error {
	gs.conn.Model(&teacher).Where("id = ?", 1).Updates(models.Teacher{FirstName: teacher.FirstName, MiddleName: teacher.MiddleName, Email: teacher.Email, Password: teacher.Password, ProfilePic: teacher.ProfilePic})
	return nil
}

func (gs *GormTeacherRepository) ReportGrade(grade models.Result, student models.Student) error {
	results := models.Result{StudentId: student.Id, Assessment: grade.Assessment, Final: grade.Final, Test: grade.Test, Total: grade.Total}
	gs.conn.Create(&results)
	return nil
}

func (gs *GormTeacherRepository) ViewClasses(room models.ClassRoom) ([]models.Student, error) {
	var class []models.Student
	gs.conn.Where("id = ?", 1).Find(&class)
	return class, nil
}
