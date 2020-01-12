package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/session"
	"net/http"
	"strings"
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

func (s *SessionUsecase) DeleteSession(id uint) []error {
	err := s.session.DeleteSession(id)
	return err
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
	if user.Role != checkUserRole(r.URL.Path) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return models.User{}, nil
	}
	return user, err
}

func (s *SessionUsecase) GetUser(id uint) (models.User, []error) {
	data, err := s.session.GetUser(id)
	return data, err
}

func checkUserRole(urlPath string) string {
	splittedPath := strings.Split(urlPath, "")
	var userRole string
	ctr := 0
	i := 0
	for {
		if ctr < 2 {
			if splittedPath[i] == "/" {
				ctr++
			} else {
				userRole += splittedPath[i]
			}
			i++
		} else if ctr == 2 {
			break
		}
	}
	return userRole
}
