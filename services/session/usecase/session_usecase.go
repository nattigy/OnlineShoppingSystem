package usecase

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	session2 "github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"net/http"
	"strconv"
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

func (s *SessionUsecase) DeleteSession(id uint, userId string, w http.ResponseWriter, r *http.Request) []error {
	cookie, _ := r.Cookie(userId)
	fmt.Println("in delete", userId)
	cookie = &http.Cookie{
		Name:     userId,
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
		Name:     strconv.Itoa(int(sess.UserID)),
		Value:    sess.Uuid,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Unix(3600, 00),
		MaxAge:   3600,
	}
	http.SetCookie(w, cookie)
	data, err := s.session.CreateSession(sess)
	return data, err
}

func (s *SessionUsecase) GetSession(value string) (models.Session, []error) {
	data, err := s.session.GetSession(value)
	return data, err
}

func (s *SessionUsecase) Check(userId string, w http.ResponseWriter, r *http.Request) (models.Session, error) {
	cookie, err := r.Cookie(userId)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return models.Session{}, err
	}
	sess, _ := s.GetSession(cookie.Value)
	if r.URL.Path == "/logout" {
		return sess, nil
	}
	if sess.Role != checkUserRole(r.URL.Path) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return models.Session{}, nil
	}
	return sess, err
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
