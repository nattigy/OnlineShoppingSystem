package usecase

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	session2 "github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"net/http"
	"strings"
	"time"
)

type SessionUsecase struct {
	session session2.SessionRepository
}

func NewSessionUsecase(session session2.SessionRepository) *SessionUsecase {
	return &SessionUsecase{session: session}
}

func (s *SessionUsecase) Sessions() ([]models.Session, []error) {
	data, err := s.session.Sessions()
	return data, err
}

func (s *SessionUsecase) DeleteSession(id uint, w http.ResponseWriter, r *http.Request) []error {
	cookie, _ := r.Cookie("session")
	cookie = &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	errs := s.session.DeleteSession(id)
	return errs
}

func (s *SessionUsecase) UpdateSession(sess models.Session) (models.Session, []error) {
	data, err := s.session.UpdateSession(sess)
	return data, err
}

func (s *SessionUsecase) CreateSession(w http.ResponseWriter, sess models.Session) (models.Session, []error) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    sess.Uuid,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	data, err := s.session.CreateSession(sess)
	return data, err
}

func (s *SessionUsecase) GetSession(value string) (models.Session, []error) {
	data, err := s.session.GetSession(value)
	return data, err
}

func (s *SessionUsecase) Check(w http.ResponseWriter, r *http.Request) (models.User, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
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
