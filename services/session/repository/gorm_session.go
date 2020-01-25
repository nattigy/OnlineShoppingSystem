package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type SessionRepository struct {
	conn *gorm.DB
}

func NewSessionRepository(conn *gorm.DB) *SessionRepository {
	return &SessionRepository{conn: conn}
}

func (s *SessionRepository) CreateSession(session models.Session) (models.Session, []error) {
	err := s.conn.Create(&session).GetErrors()
	return session, err
}

func (s *SessionRepository) GetSession(value string) (models.Session, []error) {
	var sess models.Session
	err := s.conn.Where("uuid = ?", value).First(&sess).GetErrors()
	return sess, err
}

func (s *SessionRepository) Sessions() ([]models.Session, []error) {
	var sess []models.Session
	err := s.conn.Find(&sess).GetErrors()
	return sess, err
}

func (s *SessionRepository) DeleteSession(id uint) []error {
	sess := models.Session{UserID: uint(id)}
	err := s.conn.Unscoped().Delete(&sess).GetErrors()
	return err
}

func (s *SessionRepository) UpdateSession(sess models.Session) (models.Session, []error) {
	mysession := sess
	err := s.conn.Update(&mysession).GetErrors()
	return mysession, err
}

func (s *SessionRepository) GetUser(id uint) (models.User, []error) {
	user := models.User{}
	err := s.conn.Where("id = ?", id).First(&user).GetErrors()
	user.LoggedIn = true
	return user, err
}
