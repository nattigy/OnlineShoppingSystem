package studentHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/session"
	"github.com/nattigy/parentschoolcommunicationsystem/student/usecase"
	"github.com/nattigy/parentschoolcommunicationsystem/utility"
	"html/template"
	"net/http"
	"strconv"
)

type StudentHandler struct {
	templ    *template.Template
	SUsecase usecase.StudentUsecase
	Session  session.SessionUsecase
	utility  utility.Utility
}

func NewStudentHandler(templ *template.Template, SUsecase usecase.StudentUsecase, session session.SessionUsecase, utility utility.Utility) *StudentHandler {
	return &StudentHandler{templ: templ, SUsecase: SUsecase, Session: session, utility: utility}
}

type StudentInfo struct {
	User          models.User
	Tasks         []models.Task
	UpdateProfile models.Student
	Resources     []models.Resources
	Result        []models.Result
	ClassMates    []models.Student
}

func (p *StudentHandler) ViewTasks(w http.ResponseWriter, r *http.Request) {
	user, err := p.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))

	classRoom := models.ClassRoom{
		Id: p.utility.GetSubjectById(uint(id)).ClassRoomId,
	}
	subject := models.Subject{
		Id: uint(id),
	}

	data, err := p.SUsecase.ViewTasks(classRoom, subject)
	if err != nil {
		fmt.Println(err)
	}
	in := StudentInfo{
		Tasks: data,
		User:  user,
	}
	//_ = json.NewEncoder(w).Encode(data)
	err = p.templ.ExecuteTemplate(w, "studentPortal.html", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *StudentHandler) Comment(w http.ResponseWriter, r *http.Request) {
	user, err := p.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}

	key1 := "comment"
	key2 := "taskId"

	comment := r.FormValue(key1)
	id := r.FormValue(key2)

	task := models.Task{}

	student := models.Student{
		Id: 12,
	}

	_ = p.SUsecase.Comment(task, student, comment)

	fmt.Println(id)
	http.Redirect(w, r, "/student/viewTask", http.StatusSeeOther)
}

func (p *StudentHandler) StudentUpdateProfile(w http.ResponseWriter, r *http.Request) {
	user, err := p.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}

	email := r.FormValue("studentEmail")
	password := r.FormValue("studentPassword")
	profile := r.FormValue("studentProfilePic")

	in := StudentInfo{
		User:          user,
		UpdateProfile: models.Student{Email: user.Email, Password: user.Password},
	}

	if email != "" && password != "" && profile != "" {
		user.Email = email
		user.Password = password
		in = StudentInfo{
			User:          user,
			UpdateProfile: models.Student{Email: user.Email, Password: user.Password},
		}
		studentUpdateInfo := models.Student{Email: email, Password: password, ProfilePic: profile}
		errs := p.SUsecase.StudentUpdateProfile(studentUpdateInfo)
		if errs != nil {
			fmt.Println(errs)
		}
	}
	//_ = json.NewEncoder(w).Encode(data)
	err = p.templ.ExecuteTemplate(w, "studentPortal.html", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *StudentHandler) ViewClass(w http.ResponseWriter, r *http.Request) {
	user, err := p.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}
	in := StudentInfo{
		User:       user,
		ClassMates: []models.Student{},
	}
	//_ = json.NewEncoder(w).Encode(data)
	err = p.templ.ExecuteTemplate(w, "studentPortal.html", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *StudentHandler) ViewResources(w http.ResponseWriter, r *http.Request) {
	user, err := p.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}
	in := StudentInfo{
		User:      user,
		Resources: []models.Resources{},
	}
	//_ = json.NewEncoder(w).Encode(data)
	err = p.templ.ExecuteTemplate(w, "studentPortal.html", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *StudentHandler) ViewResult(w http.ResponseWriter, r *http.Request) {
	user, err := p.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}
	in := StudentInfo{
		User:   user,
		Result: []models.Result{},
	}
	//_ = json.NewEncoder(w).Encode(data)
	err = p.templ.ExecuteTemplate(w, "studentPortal.html", in)
	if err != nil {
		fmt.Println(err)
	}
}
