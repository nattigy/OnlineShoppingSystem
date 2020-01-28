package admin

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/parentServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices"
	"net/http"
	"strconv"
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
	w.Header().Set("Content-Type", "application/json")
	var parent models.Parent
	err := json.NewDecoder(r.Body).Decode(&parent)
	fmt.Println(parent)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	er := ah.parentServices.AddParent(parent)
	if er != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
}

func (ah *AdminApi) GetParents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	parents, errs := ah.parentServices.GetParents()
	if len(errs) > 0 {
		fmt.Println(errs)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(parents)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = w.Write(response)
}

func (ah *AdminApi) GetParentById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(p.ByName("id"))
	parent, errs := ah.parentServices.GetParentById(uint(id))
	if len(errs) > 0 {
		fmt.Println(errs)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.MarshalIndent(parent, "", "\t\t")
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = w.Write(response)
}

func (ah *AdminApi) DeleteParent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(p.ByName("id"))
	errs := ah.parentServices.DeleteParent(uint(id))
	if len(errs) > 0 {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
}

func (ah *AdminApi) AddStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	fmt.Println(student)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	er := ah.studentServices.AddStudent(student)
	if er != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
}

func (ah *AdminApi) GetStudents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	students, errs := ah.studentServices.GetStudents()
	if len(errs) > 0 {
		fmt.Println(errs)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(students)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = w.Write(response)
}

func (ah *AdminApi) GetStudentById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(p.ByName("id"))
	student, errs := ah.studentServices.GetStudentById(uint(id))
	if len(errs) > 0 {
		fmt.Println(errs)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.MarshalIndent(student, "", "\t\t")
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = w.Write(response)
}

func (ah *AdminApi) DeleteStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(p.ByName("id"))
	errs := ah.studentServices.DeleteStudent(uint(id))
	if len(errs) > 0 {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
}

func (ah *AdminApi) AddTeacher(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var teacher models.Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	er := ah.teacherServices.AddTeacher(teacher)
	if er != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
}

func (ah *AdminApi) GetTeachers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	teachers, errs := ah.teacherServices.GetTeachers()
	if len(errs) > 0 {
		fmt.Println(errs)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(teachers)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = w.Write(response)
}

func (ah *AdminApi) GetTeacherById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(p.ByName("id"))
	teacher, errs := ah.teacherServices.GetTeacherById(uint(id))
	if len(errs) > 0 {
		fmt.Println(errs)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.MarshalIndent(teacher, "", "\t\t")
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = w.Write(response)
}

func (ah *AdminApi) DeleteTeacher(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(p.ByName("id"))
	errs := ah.teacherServices.DeleteTeacher(uint(id))
	if len(errs) > 0 {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
}
