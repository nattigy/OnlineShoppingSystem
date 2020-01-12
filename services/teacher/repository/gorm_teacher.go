package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type GormTeacherRepository struct {
	conn *gorm.DB
}

func NewGormTeacherRepository(Conn *gorm.DB) *GormTeacherRepository {
	return &GormTeacherRepository{conn: Conn}
}

func (gs *GormTeacherRepository) MakeNewPost(task models.Task, subject models.Subject) []error {
	task.ClassRoomId = subject.ClassRoomId
	task.SubjectId = subject.Id
	errs := gs.conn.Create(&task).GetErrors()
	return errs
}

func (gs *GormTeacherRepository) EditPost(task models.Task) []error {
	fmt.Println("task in repository", task)
	errs := gs.conn.Model(&task).Where("id = ?", task.Id).Updates(models.Task{Title: task.Title, Description: task.Description, ShortDescription: task.ShortDescription, Deadline: task.Deadline}).GetErrors()
	return errs
}

func (gs *GormTeacherRepository) RemoveTask(task models.Task) []error {
	errs := gs.conn.Unscoped().Where("id = ?", task.Id).Delete(&task).GetErrors()
	return errs
}

func (gs *GormTeacherRepository) UploadResource(resource models.Resources, s models.Subject, room models.ClassRoom) []error {
	resources := models.Resources{SubjectId: s.Id, Title: resource.Title, Description: resource.Description, Path: resource.Path}
	errs := gs.conn.Create(&resources).GetErrors()
	return errs
}

func (gs *GormTeacherRepository) TeacherUpdateProfile(teacher models.Teacher) []error {
	errs := gs.conn.Model(&teacher).Where("id = ?", teacher.Id).Updates(models.Teacher{Email: teacher.Email, Password: teacher.Password, ProfilePic: teacher.ProfilePic}).GetErrors()
	return errs
}

func (gs *GormTeacherRepository) ReportGrade(grade models.Result, student models.Student) []error {
	results := models.Result{StudentId: student.Id, Assessment: grade.Assessment, Final: grade.Final, Test: grade.Test, Total: grade.Total}
	errs := gs.conn.Create(&results).GetErrors()
	return errs
}

func (gs *GormTeacherRepository) ViewClasses(room models.ClassRoom) ([]models.Student, []error) {
	var class []models.Student
	errs := gs.conn.Where("class_room_id = ?", room.Id).Find(&class).GetErrors()
	return class, errs
}

func (gs *GormTeacherRepository) FetchPosts(subject models.Subject) ([]models.Task, []error) {
	var tasks []models.Task
	errs := gs.conn.Where("subject_id = ?", subject.Id).Find(&tasks).GetErrors()
	return tasks, errs
}
