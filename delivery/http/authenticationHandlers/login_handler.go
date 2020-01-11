package authenticationHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/parent"
	"github.com/nattigy/parentschoolcommunicationsystem/session"
	"github.com/nattigy/parentschoolcommunicationsystem/student"
	"github.com/nattigy/parentschoolcommunicationsystem/teacher"
	"github.com/nattigy/parentschoolcommunicationsystem/utility"
	"github.com/satori/uuid"
	"html/template"
	"net/http"
	"time"
)

type LoginHandler struct {
	tmpl    *template.Template
	student student.StudentUsecase
	teacher teacher.TeacherUsecase
	parent  parent.ParentUsecase
	session session.SessionUsecase
}

func NewLoginHandler(tmpl *template.Template, student student.StudentUsecase, teacher teacher.TeacherUsecase, parent parent.ParentUsecase, session session.SessionUsecase) *LoginHandler {
	return &LoginHandler{tmpl: tmpl, student: student, teacher: teacher, parent: parent, session: session}
}

func (l *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	key1 := "email"
	key2 := "password"

	user := models.User{
		Email:    r.FormValue(key1),
		Password: r.FormValue(key2),
	}

	auth, role, err := l.Authenticate(user)
	if err != nil {
		res, _ := w.Write([]byte("0"))
		fmt.Println(res, err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if auth {
		cookie, err := r.Cookie("session")
		role.LoggedIn = true
		value, _ := uuid.NewV4()

		if err != nil {
			c := &http.Cookie{
				Name:     "session",
				Value:    "",
				Path:     "/",
				MaxAge:   0,
				Expires:  time.Unix(0, 0),
				HttpOnly: true,
			}
			http.SetCookie(w, c)
		}

		cookie = &http.Cookie{
			Name:     "session",
			Value:    value.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)

		newSession := models.Session{Email: role.Email, UserID: role.Id, Uuid: cookie.Value}
		l.session.StoreSession(&newSession)

		if role.Role == utility.Student {
			http.Redirect(w, r, "/student/viewTask", http.StatusSeeOther)
		} else if role.Role == utility.Teacher {
			http.Redirect(w, r, "/teacher", http.StatusSeeOther)
		} else if role.Role == utility.Parent {
			http.Redirect(w, r, "/parent", http.StatusSeeOther)
		}
	}
}

func (l *LoginHandler) Authenticate(user models.User) (bool, models.User, error) {
	gormdb, _ := database.GormConfig()
	err := gormdb.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).GetErrors()
	if user.Id == 0 {
		return false, models.User{}, err[0]
	}
	return true, user, nil
}
