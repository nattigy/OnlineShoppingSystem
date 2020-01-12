package utility

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
)

type Utility struct {
	conn    *gorm.DB
	Session session.SessionUsecase
}

func NewUtility(conn *gorm.DB, session session.SessionUsecase) Utility {
	return Utility{conn: conn, Session: session}
}

const (
	Student string = "student"
	Teacher string = "teacher"
	Parent  string = "parent"
)

func (u *Utility) GetSubjectById(id uint) models.Subject {
	subject := models.Subject{}
	_ = u.conn.Where("id = ?", id).First(&subject).GetErrors()
	return subject
}

func (u *Utility) GetSubjectByClassRoom(grade int, section string) (models.Subject, []error) {
	subject := models.Subject{}
	classRoom := models.ClassRoom{}
	errs := u.conn.Where("grade_level = ? AND section = ?", grade, section).First(&classRoom).GetErrors()
	errs = u.conn.Where("class_room_id = ?", classRoom.Id).First(&subject).GetErrors()
	return subject, errs
}

func (u *Utility) GetSubjectByTeacherId(id uint) (models.Subject, []error) {
	subject := models.Subject{}
	errs := u.conn.Where("teacher_id = ?", id).First(&subject).GetErrors()
	return subject, errs
}

func (u *Utility) GetStudentById(id uint) (models.Student, []error) {
	student := models.Student{}
	errs := u.conn.Where("id = ?", id).First(&student).GetErrors()
	return student, errs
}

func (u *Utility) GetTeacherById(id uint) (models.Teacher, []error) {
	teacher := models.Teacher{}
	errs := u.conn.Where("id = ?", id).First(&teacher).GetErrors()
	return teacher, errs
}

func (u *Utility) GetStudentsByTeacherId(id uint) ([]models.Student, []error) {
	var students []models.Student
	student := models.Student{}
	classRooms, errs := u.GetClassRoomByTeacherId(id)
	for i := 0; i < len(classRooms); i++ {
		errs := u.conn.Where("class_room_id = ?", classRooms[i].Id).Find(&student).GetErrors()
		if errs != nil {
			fmt.Println(errs)
			//return []models.Student{}, nil
		}
		students = append(students, student)
	}

	return students, errs
}

func (u *Utility) GetClassRoomByTeacherId(id uint) ([]models.ClassRoom, []error) {
	var classRooms []models.ClassRoom
	classRoom := models.ClassRoom{}
	var subjects []models.Subject
	errs := u.conn.Where("teacher_id = ?", id).Find(&subjects).GetErrors()
	if errs != nil {
		//return classRooms, errs
	}
	fmt.Println(subjects)
	for i := 0; i < len(subjects); i++ {
		errs := u.conn.Where("id = ?", subjects[i].ClassRoomId).Find(&classRoom).GetErrors()
		if errs != nil {
			//return classRooms, errs
		}
		classRooms = append(classRooms, classRoom)
	}
	return classRooms, errs
}
