package studentApi

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices"
	"net/http"
	"strconv"
)

type StudentApi struct {
	studentService studentServices.StudentUsecase
}

func NewStudentApi(studentService studentServices.StudentUsecase) *StudentApi {
	return &StudentApi{studentService: studentService}
}

func (sa *StudentApi) ViewTasks(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))

	task, errs := sa.studentService.ViewTasks(uint(id), uint(id))
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
	_ = json.NewEncoder(w).Encode(task)
}

func (sa *StudentApi) Comment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	errs := sa.studentService.Comment(uint(id), uint(id), "", "")
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}

}

func (sa *StudentApi) ViewClass(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))

	class, errs := sa.studentService.ViewClass(uint(id))
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
	_ = json.NewEncoder(w).Encode(class)
}

func (sa *StudentApi) ViewResources(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))

	resource, errs := sa.studentService.ViewResources(uint(id))
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
	_ = json.NewEncoder(w).Encode(resource)
}

func (sa *StudentApi) ViewResult(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))

	result, errs := sa.studentService.ViewResult(uint(id))
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
	_ = json.NewEncoder(w).Encode(result)
}
