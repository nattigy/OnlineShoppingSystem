package studentApi

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices"
	"net/http"
)

type StudentApi struct {
	studentService studentServices.StudentUsecase
}

func NewStudentApi(studentService studentServices.StudentUsecase) *StudentApi {
	return &StudentApi{studentService: studentService}
}

func (sa *StudentApi) ViewTasks(w http.ResponseWriter, r *http.Request, p httprouter.Params){

}

func (sa *StudentApi) Comment(w http.ResponseWriter, r *http.Request, p httprouter.Params){

}

func (sa *StudentApi) ViewClass(w http.ResponseWriter, r *http.Request, p httprouter.Params){

}

func (sa *StudentApi) ViewResources(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (sa *StudentApi) ViewResult(w http.ResponseWriter, r *http.Request, p httprouter.Params){

}
