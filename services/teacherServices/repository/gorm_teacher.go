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

func (tr *GormTeacherRepository) AddTeacher(newTeacher models.Teacher) []error {
	errs := tr.conn.Create(&newTeacher).GetErrors()
	return errs
}

func (tr *GormTeacherRepository) GetTeachers() ([]models.Teacher, []error) {
	var teachers []models.Teacher
	errs := tr.conn.Find(&teachers).GetErrors()
	return teachers, errs
}

func (tr *GormTeacherRepository) GetTeacherById(id uint) (models.Teacher, []error) {
	teacher := models.Teacher{}
	errs := tr.conn.Where("id = ?", id).First(&teacher).GetErrors()
	return teacher, errs
}

func (tr *GormTeacherRepository) DeleteTeacher(id uint) []error {
	errs := tr.conn.Unscoped().Where("id = ?", id).Delete(&models.Teacher{}).GetErrors()
	return errs
}

func (tr *GormTeacherRepository) UpdateTeacher(newTeacher models.Teacher) (models.Teacher, []error) {
	teacher := models.Teacher{}
	errs := tr.conn.Model(&teacher).Where("id = ?", newTeacher.Id).Updates(&models.Parent{Email: newTeacher.Email, Password: newTeacher.Password}).GetErrors()
	return teacher, errs
}

func (tr *GormTeacherRepository) CreateTask(task models.Task) []error {
	errs := tr.conn.Create(&task).GetErrors()
	return errs
}

func (tr *GormTeacherRepository) GetTasks(subjectId uint) ([]models.Task, []error) {
	var tasks []models.Task
	errs := tr.conn.Where("subject_id = ?", subjectId).Find(&tasks).GetErrors()
	return tasks, errs
}

func (tr *GormTeacherRepository) UpdateTask(newTask models.Task) (models.Task, []error) {
	task := models.Task{}
	errs := tr.conn.Model(&task).Updates(&newTask).GetErrors()
	return task, errs
}

func (tr *GormTeacherRepository) DeleteTask(taskId uint) []error {
	errs := tr.conn.Unscoped().Where("id = ?", taskId).Delete(&models.Task{}).GetErrors()
	return errs
}

func (tr *GormTeacherRepository) UploadResource(resource models.Resources) []error {
	errs := tr.conn.Create(&resource).GetErrors()
	return errs
}

func (tr *GormTeacherRepository) DeleteResource(resourceId uint) []error {
	errs := tr.conn.Unscoped().Where("id = ?", resourceId).Delete(&models.Resources{}).GetErrors()
	return errs
}

func (tr *GormTeacherRepository) ReportGrade(grade models.Result) []error {
	errs := tr.conn.Create(&grade).GetErrors()
	return errs
}

func (tr *GormTeacherRepository) ViewStudents(classRoomId uint) ([]models.Student, []error) {
	var students []models.Student
	errs := tr.conn.Where("class_room_id = ?", classRoomId).Find(&students).GetErrors()
	return students, errs
}
