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

func (u *Utility) GetSubject(id uint) models.Subject {
	subject := models.Subject{}
	_ = u.conn.Where("id = ?", id).First(&subject).GetErrors()
	return subject
}
