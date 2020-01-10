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

func (s *SessionRepository) Sessions() ([]models.Session, []error) {
	var sess []models.Session
	err := s.conn.Find(&sess).GetErrors()
	return sess, err
}

func (s *SessionRepository) DeleteSession(id uint) (models.Session, []error) {
	sess := models.Session{UserID: id}
	err := s.conn.Delete(&sess).GetErrors()
	return sess, err
}

func (s *SessionRepository) UpdateSession(sess *models.Session) (*models.Session, []error) {
	mysession := sess
	err := s.conn.Update(&mysession).GetErrors()
	return mysession, err
}

func (s *SessionRepository) StoreSession(sess *models.Session) (*models.Session, []error) {
	session := sess
	err := s.conn.Create(&session).GetErrors()
	return session, err
}

func (s *SessionRepository) GetSession(value string) (models.Session, []error) {
	var sess models.Session
	err := s.conn.Where("uuid = ?").First(&sess).GetErrors()
	return sess, err
}
