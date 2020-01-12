package session

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type SessionRepository interface {
	Sessions() ([]models.Session, []error)
	DeleteSession(id uint) []error
	UpdateSession(sess models.Session) (models.Session, []error)
	StoreSession(sess models.Session) (models.Session, []error)
	GetSession(value string) (models.Session, []error)
	GetUser(id uint) (models.User, []error)
}
