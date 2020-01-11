package usecase

import (
	"fmt"
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
	data, err := s.session.Sessions()
	return data, err
}

func (s *SessionUsecase) DeleteSession(id int) (models.Session, []error) {
	data, err := s.session.DeleteSession(id)
	return data, err
}

func (s *SessionUsecase) UpdateSession(sess models.Session) (models.Session, []error) {
	data, err := s.session.UpdateSession(sess)
	return data, err
}

func (s *SessionUsecase) StoreSession(sess models.Session) (models.Session, []error) {
	data, err := s.session.StoreSession(sess)
	return data, err
}

func (s *SessionUsecase) GetSession(value string) (models.Session, []error) {
	data, err := s.session.GetSession(value)
	return data, err
}

func (s *SessionUsecase) Check(w http.ResponseWriter, r *http.Request) (models.User, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return models.User{}, err
	}
	sess, _ := s.GetSession(cookie.Value)
	user, _ := s.GetUser(sess.UserID)
	urlRole := r.URL.Path
	if user.Role != urlRole {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return models.User{}, nil
	}
	fmt.Println("user in check", user)
	return user, err
}

func (s *SessionUsecase) GetUser(id uint) (models.User, []error) {
	data, err := s.session.GetUser(id)
	return data, err
}
