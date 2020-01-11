package authenticationHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/parent"
	"github.com/nattigy/parentschoolcommunicationsystem/session"
	"github.com/nattigy/parentschoolcommunicationsystem/student"
	"github.com/nattigy/parentschoolcommunicationsystem/teacher"
	"github.com/nattigy/parentschoolcommunicationsystem/utility"
	"html/template"
	"net/http"
)

type HomePageHandler struct {
	tmpl    *template.Template
	student student.StudentUsecase
	teacher teacher.TeacherUsecase
	parent  parent.ParentUsecase
	session session.SessionUsecase
}

func NewHomePageHandler(tmpl *template.Template, student student.StudentUsecase, teacher teacher.TeacherUsecase, parent parent.ParentUsecase, session session.SessionUsecase) *HomePageHandler {
	return &HomePageHandler{tmpl: tmpl, student: student, teacher: teacher, parent: parent, session: session}
}

func (h *HomePageHandler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.tmpl.ExecuteTemplate(w, "errorPage", http.StatusSeeOther)
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		_ = h.tmpl.ExecuteTemplate(w, "index.html", false)
		return
	}
	sess, errs := h.session.GetSession(cookie.Value)
	if errs != nil {
		fmt.Println(errs)
	}
	user, errs := h.session.GetUser(sess.UserID)
	if errs != nil {
		fmt.Println(errs)
	}
	if user.Role == utility.Student {
		http.Redirect(w, r, "/student/viewTask", http.StatusSeeOther)
	} else if user.Role == utility.Teacher {
		http.Redirect(w, r, "/teacher", http.StatusSeeOther)
	} else if user.Role == utility.Parent {
		http.Redirect(w, r, "/parent", http.StatusSeeOther)
	}
}
