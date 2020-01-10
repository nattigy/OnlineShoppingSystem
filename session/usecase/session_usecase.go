package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type SessionUsecase struct {
	conn *gorm.DB
}

func NewSessionUsecase(conn *gorm.DB) *SessionUsecase {
	return &SessionUsecase{conn: conn}
}

func (s *SessionUsecase) Sessions() ([]models.Session, []error) {
	data, err := s.Sessions()
	return data, err
}

func (s *SessionUsecase) DeleteSession(id int) (models.Session, []error) {
	data, err := s.DeleteSession(id)
	return data, err
}

func (s *SessionUsecase) UpdateSession(sess *models.Session) (*models.Session, []error) {
	data, err := s.UpdateSession(sess)
	return data, err
}

func (s *SessionUsecase) StoreSession(sess *models.Session) (*models.Session, []error) {
	data, err := s.StoreSession(sess)
	return data, err
}

func (s *SessionUsecase) GetSession(value string) (models.Session, []error) {
	data, err := s.GetSession(value)
	return data, err
}

func (s *SessionUsecase) Check(sess *models.Session) (models.Session, []error) {
	data, err := s.Check(sess)
	return data, err
}
