package teacherHandlers

import (
	"fmt"
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
	w.Write([]byte("make new post"))
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
}
