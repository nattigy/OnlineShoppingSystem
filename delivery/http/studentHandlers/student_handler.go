package studentHandlers

import (
	"encoding/json"
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/student/usecase"
	"net/http"
)

type StudentHandler struct {
	SUsecase usecase.StudentUsecase
}

func NewStudentHandler(h *http.ServeMux, us usecase.StudentUsecase) {
	handler := &StudentHandler{
		SUsecase: us,
	}
	h.HandleFunc("/student/viewTask", handler.ViewTasks)
	h.HandleFunc("/student/comment", handler.Comment)
	h.HandleFunc("/student/updateProfile", handler.StudentUpdateProfile)
	h.HandleFunc("/student/viewClass", handler.ViewClass)
	h.HandleFunc("/student/viewResources", handler.ViewResources)
	h.HandleFunc("/student/viewResult", handler.ViewResult)
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
	_ = json.NewEncoder(w).Encode(data)
}

func (p *StudentHandler) Comment(w http.ResponseWriter, r *http.Request) {

}

func (p *StudentHandler) StudentUpdateProfile(w http.ResponseWriter, r *http.Request) {

}

func (p *StudentHandler) ViewClass(w http.ResponseWriter, r *http.Request) {

}

func (p *StudentHandler) ViewResources(w http.ResponseWriter, r *http.Request) {

}

func (p *StudentHandler) ViewResult(w http.ResponseWriter, r *http.Request) {

}
