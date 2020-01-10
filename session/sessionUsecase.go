package session

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type SessionUsecase interface {
	Sessions() ([]models.Session, error)
	DeleteSession(id int) (models.Session, error)
	UpdateSession(sess *models.Session) (*models.Session, error)
	StoreSession(sess *models.Session) (*models.Session, error)
	GetSession(value string) (models.Session, error)
	Check(sess *models.Session) (models.Session, error)
}
