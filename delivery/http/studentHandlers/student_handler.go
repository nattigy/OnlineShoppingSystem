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

type Info struct {
	Data []models.Task
	User models.User
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
		Id: p.utility.GetSubject(uint(id)).ClassRoomId,
	}
	subject := models.Subject{
		Id: uint(id),
	}

	data, err := p.SUsecase.ViewTasks(classRoom, subject)
	if err != nil {
		fmt.Println(err)
	}
	in := Info{
		Data: data,
		User: user,
	}
	//_ = json.NewEncoder(w).Encode(data)
	err = p.templ.ExecuteTemplate(w, "studentPortal.html", in)
	if err != nil {
		fmt.Println(err)
	}
}

func OnCardCliked(r string) string {
	fmt.Println(r)
	return "hello"
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

}
