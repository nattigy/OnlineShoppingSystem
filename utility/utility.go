package utility

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/session"
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
