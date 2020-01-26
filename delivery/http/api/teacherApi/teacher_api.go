package teacherApi

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices"
	"net/http"
)

type TeacherApi struct {
	teacherServices teacherServices.TeacherUsecase
}

func NewTeacherApi(teacherServices teacherServices.TeacherUsecase) *TeacherApi {
	return &TeacherApi{teacherServices: teacherServices}
}

func (ta *TeacherApi) CreateTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ta *TeacherApi) GetTasks(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ta *TeacherApi) UpdateTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ta *TeacherApi) DeleteTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ta *TeacherApi) UploadResource(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ta *TeacherApi) DeleteResource(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ta *TeacherApi) ReportGrade(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ta *TeacherApi) ViewStudents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
