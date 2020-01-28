package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"golang.org/x/crypto/bcrypt"
)

type GormStudentRepository struct {
	conn *gorm.DB
}

func NewGormStudentRepository(Conn *gorm.DB) *GormStudentRepository {
	return &GormStudentRepository{conn: Conn}
}

func (sr *GormStudentRepository) AddStudent(newStudent models.Student) []error {
	user := models.User{Id: newStudent.Id, Role: "student", Email: newStudent.Email, Password: newStudent.Password}
	errs := sr.conn.Create(&newStudent).GetErrors()
	errs = sr.conn.Create(&user).GetErrors()
	return errs
}

func (sr *GormStudentRepository) GetStudents() ([]models.Student, []error) {
	var students []models.Student
	errs := sr.conn.Find(&students).GetErrors()
	return students, errs
}

func (sr *GormStudentRepository) GetStudentById(id uint) (models.Student, []error) {
	student := models.Student{}
	errs := sr.conn.Where("id = ?", id).First(&student).GetErrors()
	return student, errs
}

func (sr *GormStudentRepository) DeleteStudent(id uint) []error {
	errs := sr.conn.Unscoped().Where("id = ?", id).Delete(&models.Student{}).GetErrors()
	errs = sr.conn.Unscoped().Where("id = ?", id).Delete(&models.User{}).GetErrors()
	return errs
}

func (sr *GormStudentRepository) UpdateStudent(newStudent models.Student) (models.Student, []error) {
	student := models.Student{}
	user := models.User{}
	password, _ := bcrypt.GenerateFromPassword([]byte(newStudent.Password), bcrypt.DefaultCost)
	errs := sr.conn.Model(&student).Where("id = ?", newStudent.Id).Updates(&models.Student{Email: newStudent.Email, Password: string(password)}).GetErrors()
	errs = sr.conn.Model(&user).Where("id = ?", newStudent.Id).Updates(&models.User{Email: newStudent.Email, Password: string(password)}).GetErrors()
	return student, errs
}

func (sr *GormStudentRepository) ViewTasks(classRoomId uint, subjectId uint) ([]models.Task, []error) {
	var tasks []models.Task
	errs := sr.conn.Where("class_room_id = ? AND subject_id = ?", classRoomId, subjectId).FirstOrInit(&tasks).GetErrors()
	return tasks, errs
}

func (sr *GormStudentRepository) Comment(taskId uint, studentId uint, studentName string, data string) []error {
	errs := sr.conn.Create(&models.Comment{TaskId: taskId, StudentId: studentId, FirstName: studentName, Data: data}).GetErrors()
	return errs
}

func (sr *GormStudentRepository) ViewClass(sectionId uint) ([]models.Student, []error) {
	var students []models.Student
	errs := sr.conn.Where("section_id = ?", sectionId).Find(&students).GetErrors()
	return students, errs
}

func (sr *GormStudentRepository) Comments(taskId uint) ([]models.Comment, []error) {
	var comments []models.Comment
	errs := sr.conn.Where("task_id = ?", taskId).Find(&comments).GetErrors()
	return comments, errs
}

func (sr *GormStudentRepository) ViewResources(subjectId uint) ([]models.Resources, []error) {
	var resources []models.Resources
	errs := sr.conn.Where("subject_id = ?", subjectId).Find(&resources).GetErrors()
	return resources, errs
}

func (sr *GormStudentRepository) ViewResult(studentId uint) (models.Student, []error) {
	student := models.Student{}
	var result []models.Result
	errs := sr.conn.Where("student_id = ?", studentId).Find(&result).GetErrors()
	errs = sr.conn.Where("id = ?", studentId).Find(&student).GetErrors()
	student.Result = result
	return student, errs
}

func (sr *GormStudentRepository) GetHomeRoomTeacher(studentId uint) (models.Teacher, []error) {
	teacher := models.Teacher{}
	student, errs := sr.GetStudentById(studentId)
	errs = sr.conn.Where("class_room_id = ?", student.ClassRoomId).First(&teacher).GetErrors()
	return teacher, errs
}
