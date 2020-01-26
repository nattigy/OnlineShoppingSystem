package authenticationHandlers

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/uuid"
	"html/template"
	"net/http"

	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/parentServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/utility"
)

const (
	Student string = "student"
	Teacher string = "teacher"
	Parent  string = "parent"
)

type AuthenticationHandler struct {
	tmpl    *template.Template
	student studentServices.StudentUsecase
	teacher teacherServices.TeacherUsecase
	parent  parentServices.ParentUsecase
	session session.SessionUsecase
	utility utility.AuthenticationUsecase
}

func NewAuthenticationHandler(tmpl *template.Template, student studentServices.StudentUsecase, teacher teacherServices.TeacherUsecase, parent parentServices.ParentUsecase, session session.SessionUsecase, utility utility.AuthenticationUsecase) *AuthenticationHandler {
	return &AuthenticationHandler{tmpl: tmpl, student: student, teacher: teacher, parent: parent, session: session, utility: utility}
}

func (l *AuthenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	cookie, er := r.Cookie("session")
	if er == nil {
		sessId, _ := l.session.GetSession(cookie.Value)
		_ = l.session.DeleteSession(sessId.ID, w, r)
	}

	auth, role, err := l.utility.Authenticate(user)
	if err != nil {
		fmt.Println("login err : ", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if auth {
		role.LoggedIn = true
		value, _ := uuid.NewV4()
		ses := models.Session{Role: role.Role, Email: role.Email, UserID: role.Id, Uuid: value.String()}
		l.session.CreateSession(w, ses)
		if er == nil {
			u, err2 := l.session.GetSession(value.String())
			if err2 != nil {
				fmt.Println(err2)
			}
			Redirect(w, r, models.User{Role: u.Role})
			return
		}
		Redirect(w, r, role)
		return
	}

}

func (l *AuthenticationHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	sessId, _ := l.session.GetSession(cookie.Value)
	errors := l.session.DeleteSession(sessId.ID, w, r)
	if errors != nil {
		fmt.Println(errors)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (l *AuthenticationHandler) AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := l.session.Check(w, r)
		ctx := context.WithValue(r.Context(), "signed_in_user_session", user)
		next.ServeHTTP(w, r.WithContext(ctx))
		if err != nil {
			fmt.Println(err)
			return
		}
		if user.Id == 0 {
			fmt.Println("Id not found")
			return
		}
	})
}

func (l *AuthenticationHandler) UserHandler(next httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		user, err := l.session.Check(w, r)
		if err != nil {
			fmt.Println(err)
			return
		}
		if user.Id == 0 {
			fmt.Println("Id not found")
			return
		}
	})
}

func Redirect(w http.ResponseWriter, r *http.Request, role models.User) {
	if role.Role == Student {
		http.Redirect(w, r, "/student/viewTask?subjectId=100", http.StatusSeeOther)
	} else if role.Role == Teacher {
		http.Redirect(w, r, "/teacher/makeNewPost", http.StatusSeeOther)
	} else if role.Role == Parent {
		http.Redirect(w, r, "/parent/viewGrade", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
