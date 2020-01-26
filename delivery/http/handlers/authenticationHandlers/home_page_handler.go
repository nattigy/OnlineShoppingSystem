package authenticationHandlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/nattigy/parentschoolcommunicationsystem/services/parentServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices"
)

type HomePageHandler struct {
	tmpl    *template.Template
	student studentServices.StudentUsecase
	teacher teacherServices.TeacherUsecase
	parent  parentServices.ParentUsecase
	session session.SessionUsecase
}

func NewHomePageHandler(tmpl *template.Template, student studentServices.StudentUsecase, teacher teacherServices.TeacherUsecase, parent parentServices.ParentUsecase, session session.SessionUsecase) *HomePageHandler {
	return &HomePageHandler{tmpl: tmpl, student: student, teacher: teacher, parent: parent, session: session}
}

func (h *HomePageHandler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		_ = h.tmpl.ExecuteTemplate(w, "errorPage", http.StatusSeeOther)
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		_ = h.tmpl.ExecuteTemplate(w, "home.layout", false)
		return
	}
	sess, errs := h.session.GetSession(cookie.Value)
	if errs != nil {
		fmt.Println(errs)
	}
	user, errs := h.session.GetUser(sess.UserID)
	Redirect(w, r, user)
}
