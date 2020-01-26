package admin

import (
	"encoding/json"
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
	FirstName := r.FormValue("firstname")
	MiddleName := r.FormValue("middlename")
	id, _ := strconv.Atoi(p.ByName("id"))
	Email := r.FormValue("email")
	Password := r.FormValue("password")
	ProfilePic := r.FormValue("profilePic")
	errs := ah.parentServices.AddParent(models.Parent{FirstName: FirstName, MiddleName: MiddleName, Id: uint(id), Email: Email, Password: Password, ProfilePic: ProfilePic})
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
}

func (ah *AdminApi) GetParents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	parents, errs := ah.parentServices.GetParents()
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
	_ = json.NewEncoder(w).Encode(parents)
}

func (ah *AdminApi) GetParentById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))

	parents, errs := ah.parentServices.GetParentById(uint(id))
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
	_ = json.NewEncoder(w).Encode(parents)
}

func (ah *AdminApi) DeleteParent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	errs := ah.parentServices.DeleteParent(uint(id))
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
	_ = json.NewEncoder(w).Encode(errs)

}

func (ah *AdminApi) AddStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	FirstName := r.FormValue("firstname")
	MiddleName := r.FormValue("middlename")
	Email := r.FormValue("email")
	Password := r.FormValue("password")
	SectionId, _ := strconv.Atoi(r.FormValue("sectionid"))
	ClassRoomId, _ := strconv.Atoi(r.FormValue("classroomid"))
	ParentId, _ := strconv.Atoi(r.FormValue("parentid"))
	ProfilePic := r.FormValue("profilepic")

	errs := ah.studentServices.AddStudent(models.Student{FirstName: FirstName, MiddleName: MiddleName, Id: uint(id), Email: Email, Password: Password, ProfilePic: ProfilePic, SectionId: uint(SectionId), ClassRoomId: uint(ClassRoomId), ParentId: uint(ParentId)})
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
}

func (ah *AdminApi) GetStudents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	students, errs := ah.studentServices.GetStudents()
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
	_ = json.NewEncoder(w).Encode(students)
}

func (ah *AdminApi) GetStudentById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))

	student, errs := ah.studentServices.GetStudentById(uint(id))
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
	_ = json.NewEncoder(w).Encode(student)
}

func (ah *AdminApi) DeleteStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	errs := ah.studentServices.DeleteStudent(uint(id))
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
}

func (ah *AdminApi) AddTeacher(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	FirstName := r.FormValue("firstname")
	MiddleName := r.FormValue("middlename")
	Email := r.FormValue("email")
	Password := r.FormValue("password")
	ProfilePic := r.FormValue("profilepic")
	SubjectId, _ := strconv.Atoi(r.FormValue("subjectid"))
	ClassRoomId, _ := strconv.Atoi(r.FormValue("classRoomid"))

	errs := ah.teacherServices.AddTeacher(models.Teacher{FirstName: FirstName, MiddleName: MiddleName, Id: uint(id), Email: Email, Password: Password, ProfilePic: ProfilePic, ClassRoomId: uint(ClassRoomId), SubjectId: uint(SubjectId)})
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
}

func (ah *AdminApi) GetTeachers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	teachers, errs := ah.teacherServices.GetTeachers()
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
	_ = json.NewEncoder(w).Encode(teachers)
}

func (ah *AdminApi) GetTeacherById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))

	teacher, errs := ah.teacherServices.GetTeacherById(uint(id))
	if len(errs) > 0 {
		_ = json.NewEncoder(w).Encode(errs)
	}
	_ = json.NewEncoder(w).Encode(teacher)
}

func (ah *AdminApi) DeleteTeacher(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	errs := ah.teacherServices.DeleteTeacher(uint(id))
	if len(errs) > 0 {

		_ = json.NewEncoder(w).Encode(errs)
	}
}
