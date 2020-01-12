package session

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"net/http"
)

type SessionUsecase interface {
	Sessions() ([]models.Session, []error)
	DeleteSession(id uint) []error
	UpdateSession(sess models.Session) (models.Session, []error)
	StoreSession(sess models.Session) (models.Session, []error)
	GetSession(value string) (models.Session, []error)
	Check(w http.ResponseWriter, r *http.Request) (models.User, error)
	GetUser(id uint) (models.User, []error)
}
