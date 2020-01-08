package studentHandlers

import (
	"encoding/json"
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/authenticate"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/student/usecase"
	"html/template"
	"net/http"
	"strconv"
)

type StudentHandler struct {
	templ    *template.Template
	SUsecase usecase.StudentUsecase
}

func NewStudentHandler(t *template.Template, us usecase.StudentUsecase) *StudentHandler {
	return &StudentHandler{
		templ:    t,
		SUsecase: us,
	}
}

type Info struct {
	Data []models.Task
	User authenticate.User
}

func (p *StudentHandler) ViewTasks(w http.ResponseWriter, r *http.Request) {
	classRoom := models.ClassRoom{
		Id:         12,
		GradeLevel: 12,
		Section:    "a",
	}
	id, _ := strconv.Atoi(r.FormValue("id"))
	subject := models.Subject{
		Id: id,
	}
	role := authenticate.Role{
		Student: true,
		Teacher: false,
		Parent:  false,
	}

	user := authenticate.User{
		Role:     role,
		Loggedin: true,
	}
	var stu models.Student
	cookie, _ := r.Cookie("session")
	bb := []byte(cookie.Value)

	fmt.Println(string(bb))
	err := json.Unmarshal(bb, &stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu)

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

	key1 := "comment"
	key2 := "taskId"

	comment := r.FormValue(key1)
	id := r.FormValue(key2)

	task := models.Task{
		Id: 12,
	}

	student := models.Student{
		Id: 12,
	}

	_ = p.SUsecase.Comment(task, student, comment)

	fmt.Println(id)
	http.Redirect(w, r, "/student/viewTask", http.StatusSeeOther)
}

func (p *StudentHandler) StudentUpdateProfile(w http.ResponseWriter, r *http.Request) {

}

func (p *StudentHandler) ViewClass(w http.ResponseWriter, r *http.Request) {

}

func (p *StudentHandler) ViewResources(w http.ResponseWriter, r *http.Request) {

}

func (p *StudentHandler) ViewResult(w http.ResponseWriter, r *http.Request) {

}
