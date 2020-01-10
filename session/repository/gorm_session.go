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

func (s *SessionRepository) Sessions() ([]models.Session, error) {
	panic("implement me")
}

func (s *SessionRepository) DeleteSession(id int) (models.Session, error) {
	panic("implement me")
}

func (s *SessionRepository) UpdateSession(sess *models.Session) (*models.Session, error) {
	panic("implement me")
}

func (s *SessionRepository) StoreSession(sess *models.Session) (*models.Session, error) {
	panic("implement me")
}

func (s *SessionRepository) GetSession(value string) (models.Session, error) {
	panic("implement me")
}
