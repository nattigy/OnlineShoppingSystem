package teacherHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacher/usecase"
	"github.com/nattigy/parentschoolcommunicationsystem/services/utility"
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
	Resource      models.Resources
	UpdateProfile models.Teacher
	Students      []models.Student
	FetchPost     []models.Task
	Task          models.Task
	Result        []models.Result
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
	}

	err = t.templ.ExecuteTemplate(w, "teacherMakeNewPost", in)
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
		User: user,
	}
	err = t.templ.ExecuteTemplate(w, "teacherUploadResource", in)
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
	teacher, _ := t.utility.GetTeacherById(user.Id)
	in := TeacherInfo{
		User:          user,
		UpdateProfile: teacher,
	}
	err = t.templ.ExecuteTemplate(w, "teacherUpdateProfile", in)
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
		User: user,
	}
	err = t.templ.ExecuteTemplate(w, "teacherReportGrade", in)
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
	students, errs := t.utility.GetStudentsByTeacherId(user.Id)
	if errs != nil {
		fmt.Println(errs)
	}
	in := TeacherInfo{
		User:     user,
		Students: students,
	}
	err = t.templ.ExecuteTemplate(w, "teacherViewClasses", in)
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
		FetchPost: prevoiusPosts,
		Task:      prevoiusPosts[0],
	}
	err = t.templ.ExecuteTemplate(w, "teacherEditPost", in)
	if err != nil {
		fmt.Println(err)
	}
}
