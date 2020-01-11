package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/session"
	"net/http"
)

type SessionUsecase struct {
	session session.SessionRepository
}

func NewSessionUsecase(session session.SessionRepository) *SessionUsecase {
	return &SessionUsecase{session: session}
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

func (s *SessionUsecase) Check(w http.ResponseWriter, r *http.Request) (models.User, error) {
	cookie, err := r.Cookie("session")
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return models.User{}, nil
	}
	sess, _ := s.GetSession(cookie.Value)
	user, _ := s.GetUser(sess.UserID)
	if user.Role != sess.Role {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	return user, err
}

func (s *SessionUsecase) GetUser(id uint) (models.User, []error) {
	data, err := s.GetUser(id)
	return data, err
}
