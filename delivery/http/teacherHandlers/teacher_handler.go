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
	subject, errs := t.utility.GetSubjectByClassRoom(grade, section)

	if errs != nil {
		fmt.Println(errs)
	}

	if title != "" && shortDescriptio != "" && description != "" && grade != 0 && section != "" {
		errs = t.TUsecase.MakeNewPost(newTask, subject)
		if errs != nil {
			fmt.Println(errs)
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
	editTitle := r.FormValue("editTitle")
	editDate := r.FormValue("editDate")
	editDescription := r.FormValue("editDescription")
	id, _ := strconv.Atoi(r.FormValue("id"))

	editedTask := models.Task{Id: uint(id), Title: editTitle, Deadline: editDate, ShortDescription: editDescription}

	errs := t.TUsecase.EditPost(editedTask)
	if errs != nil {
		fmt.Println(errs)
	}
	http.Redirect(w, r, "/teacher/fetchPosts", http.StatusSeeOther)
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
	id, err := strconv.Atoi(r.FormValue("id"))
	errs := t.TUsecase.RemoveTask(models.Task{Id: uint(id)})
	if errs != nil {
		fmt.Println(errs)
	}
	http.Redirect(w, r, "/teacher/fetchPosts", http.StatusSeeOther)
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

func (t *TeacherHandler) FetchPosts(w http.ResponseWriter, r *http.Request) {
	user, err := t.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}
	subject, _ := t.utility.GetSubjectByTeacherId(user.Id)
	prevoiusPosts, errs := t.TUsecase.FetchPosts(subject)
	if errs != nil {
		fmt.Println(errs)
	}
	in := TeacherInfo{
		User:      user,
		FetchPost: true,
		Data:      data{FetchPost: prevoiusPosts},
	}
	err = t.templ.ExecuteTemplate(w, "teacherPortal.html", in)
	if err != nil {
		fmt.Println(err)
	}
}
