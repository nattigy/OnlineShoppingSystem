package authenticationHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/parent"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"github.com/nattigy/parentschoolcommunicationsystem/services/student"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacher"
	"github.com/nattigy/parentschoolcommunicationsystem/services/utility"
	"github.com/satori/uuid"
	"html/template"
	"net/http"
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if auth {
		cookie, err := r.Cookie("session")
		role.LoggedIn = true
		value, _ := uuid.NewV4()

		if err == nil {
			u, err := l.session.GetSession(cookie.Value)
			if u.Role == utility.Student {
				http.Redirect(w, r, "/student/viewTask?id=1", http.StatusSeeOther)
				return
			} else if u.Role == utility.Teacher {
				http.Redirect(w, r, "/teacher", http.StatusSeeOther)
				return
			} else if u.Role == utility.Parent {
				http.Redirect(w, r, "/parent", http.StatusSeeOther)
				return
			}
			if err != nil {

			}
		}

		cookie = &http.Cookie{
			Name:     "session",
			Value:    value.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)

		newSession := models.Session{Email: role.Email, UserID: role.Id, Uuid: cookie.Value, Role: role.Role}
		l.session.StoreSession(newSession)

		if role.Role == utility.Student {
			fmt.Println("here student")
			http.Redirect(w, r, "/student/viewTask?id=1", http.StatusSeeOther)
		} else if role.Role == utility.Teacher {
			fmt.Println("here teacher")
			http.Redirect(w, r, "/teacher/makeNewPost", http.StatusSeeOther)
		} else if role.Role == utility.Parent {
			fmt.Println("here parent")
			http.Redirect(w, r, "/parent/viewGrade", http.StatusSeeOther)
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
