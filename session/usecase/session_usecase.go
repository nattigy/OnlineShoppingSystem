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

func (s *SessionUsecase) Sessions() ([]models.Session, error) {
	panic("implement me")
}

func (s *SessionUsecase) DeleteSession(id int) (models.Session, error) {
	panic("implement me")
}

func (s *SessionUsecase) UpdateSession(sess *models.Session) (*models.Session, error) {
	panic("implement me")
}

func (s *SessionUsecase) StoreSession(sess *models.Session) (*models.Session, error) {
	panic("implement me")
}

func (s *SessionUsecase) GetSession(value string) (models.Session, error) {
	panic("implement me")
}

func (s *SessionUsecase) Check(sess *models.Session) (models.Session, error) {
	panic("implement me")
}
