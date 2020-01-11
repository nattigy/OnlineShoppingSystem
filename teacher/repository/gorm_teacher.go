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

func (gs *GormTeacherRepository) MakeNewPost(task models.Task, subject models.Subject) error {
	task.ClassRoomId = subject.ClassRoomId
	task.SubjectId = subject.Id
	errs := gs.conn.Create(&task).GetErrors()
	return errs[0]
}

func (gs *GormTeacherRepository) EditPost(task models.Task) error {
	errs := gs.conn.Model(&task).Where("id = ?", task.Id).Updates(models.Task{Title: task.Title, Description: task.Description, ShortDescription: task.ShortDescription, Deadline: task.Deadline}).GetErrors()
	return errs[0]
}

func (gs *GormTeacherRepository) RemoveTask(task models.Task) error {
	errs := gs.conn.Where("id = ?", task.Id).Delete(&task).GetErrors()
	return errs[0]
}

func (gs *GormTeacherRepository) UploadResource(resource models.Resources, s models.Subject, room models.ClassRoom) error {
	resources := models.Resources{SubjectId: s.Id, Title: resource.Title, Description: resource.Description, Path: resource.Path}
	errs := gs.conn.Create(&resources).GetErrors()
	return errs[0]
}

func (gs *GormTeacherRepository) TeacherUpdateProfile(teacher models.Teacher) error {
	errs := gs.conn.Model(&teacher).Where("id = ?", teacher.Id).Updates(models.Teacher{Email: teacher.Email, Password: teacher.Password, ProfilePic: teacher.ProfilePic}).GetErrors()
	return errs[0]
}

func (gs *GormTeacherRepository) ReportGrade(grade models.Result, student models.Student) error {
	results := models.Result{StudentId: student.Id, Assessment: grade.Assessment, Final: grade.Final, Test: grade.Test, Total: grade.Total}
	errs := gs.conn.Create(&results).GetErrors()
	return errs[0]
}

func (gs *GormTeacherRepository) ViewClasses(room models.ClassRoom) ([]models.Student, error) {
	var class []models.Student
	gs.conn.Where("class_room_id = ?", room.Id).Find(&class)
	return class, nil
}
