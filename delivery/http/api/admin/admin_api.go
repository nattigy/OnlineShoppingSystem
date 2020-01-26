package admin

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/parentServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices"
	"net/http"
)

type AdminApi struct {
	studentServices studentServices.StudentUsecase
	teacherServices teacherServices.TeacherUsecase
	parentServices  parentServices.ParentUsecase
}

func NewAdminApi(studentServices studentServices.StudentUsecase, teacherServices teacherServices.TeacherUsecase, parentServices parentServices.ParentUsecase) *AdminApi {
	return &AdminApi{studentServices: studentServices, teacherServices: teacherServices, parentServices: parentServices}
}

func (ah *AdminApi) AddParent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	parentName := r.FormValue("firstname")
	errs := ah.parentServices.AddParent(models.Parent{FirstName: parentName})
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode("")
	}
}

func (ah *AdminApi) GetParents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	parents, errs := ah.parentServices.GetParents()
	if len(errs) > 0 {

	}
	_ = json.NewEncoder(w).Encode(parents)
}

func (ah *AdminApi) GetParentById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ah *AdminApi) DeleteParent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ah *AdminApi) AddStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ah *AdminApi) GetStudents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ah *AdminApi) GetStudentById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ah *AdminApi) DeleteStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ah *AdminApi) AddTeacher(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ah *AdminApi) GetTeachers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ah *AdminApi) GetTeacherById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (ah *AdminApi) DeleteTeacher(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
