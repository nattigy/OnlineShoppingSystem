package teacherHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/session"
	"github.com/nattigy/parentschoolcommunicationsystem/teacher/usecase"
	"github.com/nattigy/parentschoolcommunicationsystem/utility"
	"html/template"
	"net/http"
)

type TeacherHandler struct {
	templ    *template.Template
	TUsecase usecase.TeacherUsecase
	Session  session.SessionUsecase
	utility  utility.Utility
}

func NewTeacherHandler(templ *template.Template, TUsecase usecase.TeacherUsecase, session session.SessionUsecase, utility utility.Utility) *TeacherHandler {
	return &TeacherHandler{templ: templ, TUsecase: TUsecase, Session: session, utility: utility}
}

type TeacherInfo struct {
	User          models.User
	Post          models.Task
	Resource      models.Resources
	UpdateProfile models.Teacher
}

func (t *TeacherHandler) MakeNewPost(w http.ResponseWriter, r *http.Request) {
	user, err := t.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}
	in := TeacherInfo{
		User: user,
	}
	t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
}

func (t *TeacherHandler) EditPost(w http.ResponseWriter, r *http.Request) {
	user, err := t.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}
	in := TeacherInfo{
		User: user,
	}
	t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
}

func (t *TeacherHandler) RemoveTask(w http.ResponseWriter, r *http.Request) {
	user, err := t.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}
	in := TeacherInfo{
		User: user,
	}
	t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
}

func (t *TeacherHandler) UploadResource(w http.ResponseWriter, r *http.Request) {
	user, err := t.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}
	in := TeacherInfo{
		User: user,
	}
	t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
}

func (t *TeacherHandler) TeacherUpdateProfile(w http.ResponseWriter, r *http.Request) {
	user, err := t.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}
	in := TeacherInfo{
		User: user,
	}
	t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
}

func (t *TeacherHandler) ReportGrade(w http.ResponseWriter, r *http.Request) {
	user, err := t.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}
	in := TeacherInfo{
		User: user,
	}
	t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
}

func (t *TeacherHandler) ViewClasses(w http.ResponseWriter, r *http.Request) {
	user, err := t.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}
	in := TeacherInfo{
		User: user,
	}
	t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
}
