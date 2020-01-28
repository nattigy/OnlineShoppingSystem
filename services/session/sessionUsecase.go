package session

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"net/http"
)

type SessionUsecase interface {
	Sessions() ([]models.Session, []error)
	DeleteSession(id uint, userId string, w http.ResponseWriter, r *http.Request) []error
	UpdateSession(sess models.Session) (models.Session, []error)
	CreateSession(w http.ResponseWriter, sess models.Session) (models.Session, []error)
	GetSession(value string) (models.Session, []error)
	Check(userId string, w http.ResponseWriter, r *http.Request) (models.Session, error)
	GetUser(id uint) (models.User, []error)
}
