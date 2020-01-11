package teacherHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/session"
	"github.com/nattigy/parentschoolcommunicationsystem/teacher/usecase"
	"github.com/nattigy/parentschoolcommunicationsystem/utility"
	"html/template"
	"net/http"
	"strconv"
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
	Post          bool
	Edit          bool
	Resource      bool
	UpdateProfile bool
	Students      bool
	Data          data
	FetchPost     bool
	UploadResult  bool
}

type data struct {
	Resource      models.Resources
	UpdateProfile models.Teacher
	Students      []models.Student
	FetchPost     []models.Task
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

	title := r.FormValue("title")
	shortDescriptio := r.FormValue("shortDescription")
	description := r.FormValue("description")
	grade, _ := strconv.Atoi(r.FormValue("grade"))
	section := r.FormValue("section")

	newTask := models.Task{Title: title, ShortDescription: shortDescriptio, Description: description}
	subject, err := t.utility.GetSubjectByClassRoom(grade, section)

	if err != nil {
		fmt.Println(err)
	}

	if title != "" && shortDescriptio != "" && description != "" && grade != 0 && section != "" {
		err = t.TUsecase.MakeNewPost(newTask, subject)
		if err != nil {
			fmt.Println(err)
		}
	}

	in := TeacherInfo{
		User: user,
		Post: true,
	}

	err = t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
	if err != nil {
		fmt.Println(err)
	}
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
		Edit: true,
	}
	err = t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
	if err != nil {
		fmt.Println(err)
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
	in := TeacherInfo{
		User: user,
	}
	err = t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
	if err != nil {
		fmt.Println(err)
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
	in := TeacherInfo{
		User:     user,
		Resource: true,
		Data:     data{Resource: models.Resources{}},
	}
	err = t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
	if err != nil {
		fmt.Println(err)
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
	in := TeacherInfo{
		User:          user,
		UpdateProfile: true,
		Data:          data{UpdateProfile: models.Teacher{}},
	}
	err = t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
	if err != nil {
		fmt.Println(err)
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
	in := TeacherInfo{
		User:         user,
		UploadResult: true,
	}
	err = t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
	if err != nil {
		fmt.Println(err)
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
	in := TeacherInfo{
		User:     user,
		Students: true,
		Data:     data{Students: []models.Student{}},
	}
	err = t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
	if err != nil {
		fmt.Println(err)
	}
}
