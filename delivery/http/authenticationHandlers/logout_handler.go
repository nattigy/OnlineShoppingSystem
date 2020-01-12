package authenticationHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/parent"
	"github.com/nattigy/parentschoolcommunicationsystem/session"
	"github.com/nattigy/parentschoolcommunicationsystem/student"
	"github.com/nattigy/parentschoolcommunicationsystem/teacher"
	"html/template"
	"net/http"
	"time"
)

type LogoutHandler struct {
	tmpl    *template.Template
	student student.StudentUsecase
	teacher teacher.TeacherUsecase
	parent  parent.ParentUsecase
	session session.SessionUsecase
}

func NewLogoutHandler(tmpl *template.Template, student student.StudentUsecase, teacher teacher.TeacherUsecase, parent parent.ParentUsecase, session session.SessionUsecase) *LogoutHandler {
	return &LogoutHandler{tmpl: tmpl, student: student, teacher: teacher, parent: parent, session: session}
}

func (l *LogoutHandler) Logout(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, "/", http.StatusSeeOther)
	sessId, _ := l.session.GetSession(cookie.Value)
	errors := l.session.DeleteSession(sessId.ID)
	if errors != nil {
		fmt.Println(errors)
	}
}
