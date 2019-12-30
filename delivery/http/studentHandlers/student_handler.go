package studentHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/student/usecase"
	"html/template"
	"net/http"
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

func (p *StudentHandler) ViewTasks(w http.ResponseWriter, r *http.Request) {
	classRoom := models.ClassRoom{
		Id:         12,
		GradeLevel: 12,
		Section:    "a",
	}
	subject := models.Subject{
		Id: 31,
	}
	data, err := p.SUsecase.ViewTasks(classRoom, subject)
	if err != nil {
		fmt.Println(err)
	}
	//_ = json.NewEncoder(w).Encode(data)
	_ = p.templ.ExecuteTemplate(w, "studentPortal.html", data)
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
	http.RedirectHandler("/student/viewTask", 200)
}

func (p *StudentHandler) StudentUpdateProfile(w http.ResponseWriter, r *http.Request) {

}

func (p *StudentHandler) ViewClass(w http.ResponseWriter, r *http.Request) {

}

func (p *StudentHandler) ViewResources(w http.ResponseWriter, r *http.Request) {

}

func (p *StudentHandler) ViewResult(w http.ResponseWriter, r *http.Request) {

}
