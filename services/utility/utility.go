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
		if len(errs) != 0 {
			fmt.Println(errs)
			return []models.Student{}, errs
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
	if len(errs) != 0 {
		fmt.Println(errs)
		return classRooms, errs
	}
	for i := 0; i < len(subjects); i++ {
		errs := u.conn.Where("id = ?", subjects[i].ClassRoomId).Find(&classRoom).GetErrors()
		if len(errs) != 0 {
			fmt.Println(errs)
			return classRooms, errs
		}
		classRooms = append(classRooms, classRoom)
	}
	return classRooms, errs
}

func (u *Utility) GetTeacherByParentId(id uint) (models.Teacher, []error) {
	student := models.Student{}
	errs := u.conn.Where("parent_id = ?", id).First(&student).GetErrors()
	if len(errs) != 0 {
		fmt.Println(errs)
		return models.Teacher{}, errs
	}
	classRoom := models.ClassRoom{}
	errs = u.conn.Where("id = ?", student.ClassRoomId).First(&classRoom).GetErrors()
	if len(errs) != 0 {
		fmt.Println(errs)
		return models.Teacher{}, errs
	}
	return models.Teacher{Id: classRoom.HomeRoom}, nil
}

func (u *Utility) GetParentsByTeacherId(id uint) ([]models.Parent, []error) {
	classRoom := models.ClassRoom{}
	errs := u.conn.Where("home_room = ?", id).First(&classRoom).GetErrors()
	if len(errs) != 0 {
		fmt.Println(errs)
		return []models.Parent{}, errs
	}
	var students []models.Student
	errs = u.conn.Where("class_room_id = ?", classRoom.Id).Find(&students).GetErrors()
	parent := models.Parent{}
	var parents []models.Parent
	for i := 0; i < len(students); i++ {
		errs = u.conn.Where("id = ?", students[i].ParentId).Find(&parent).GetErrors()
		if len(errs) != 0 {
			fmt.Println(errs)
			return nil, errs
		}
		parents = append(parents, parent)
	}
	return parents, nil
}
