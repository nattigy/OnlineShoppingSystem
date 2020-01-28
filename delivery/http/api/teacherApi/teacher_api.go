package teacherApi

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices"
	"net/http"
	"strconv"
)

type TeacherApi struct {
	teacherServices teacherServices.TeacherUsecase
}

func NewTeacherApi(teacherServices teacherServices.TeacherUsecase) *TeacherApi {
	return &TeacherApi{teacherServices: teacherServices}
}

func (ta *TeacherApi) CreateTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))

	//CreatedAt        time.Time
	//UpdatedAt        time.Time
	//DeletedAt        *time.Time `sql:"index"`
	title := r.FormValue("title")
	description := r.FormValue("description")
	shortDescription := r.FormValue("shortdescription")
	subjectId, _ := strconv.Atoi(r.FormValue("subjectid"))
	classRoomId, _ := strconv.Atoi(r.FormValue("classroomid"))
	deadline := r.FormValue("deadline")

	errs := ta.teacherServices.CreateTask(models.Task{Id: uint(id), Title: title, Description: description, ShortDescription: shortDescription, SubjectId: uint(subjectId), ClassRoomId: uint(classRoomId), Deadline: deadline})
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	http.StatusText(http.StatusOK)

}

func (ta *TeacherApi) GetTasks(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	tasks, errs := ta.teacherServices.GetTasks(uint(id))
	output, _ := json.MarshalIndent(tasks, "", "\t\t")
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

func (ta *TeacherApi) UpdateTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	title := r.FormValue("title")
	description := r.FormValue("description")
	shortDescription := r.FormValue("shortdescription")
	subjectId, _ := strconv.Atoi(r.FormValue("subjectid"))
	classRoomId, _ := strconv.Atoi(r.FormValue("classroomid"))
	deadline := r.FormValue("deadline")

	tasks, errs := ta.teacherServices.UpdateTask(models.Task{Id: uint(id), Title: title, Description: description, ShortDescription: shortDescription, SubjectId: uint(subjectId), ClassRoomId: uint(classRoomId), Deadline: deadline})
	output, _ := json.MarshalIndent(tasks, "", "\t\t")
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

func (ta *TeacherApi) DeleteTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	errs := ta.teacherServices.DeleteTeacher(uint(id))
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	http.StatusText(http.StatusOK)
}

func (ta *TeacherApi) UploadResource(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	subjectId, _ := strconv.Atoi(r.FormValue("subjectid"))
	title := r.FormValue("title")
	description := r.FormValue("description")
	path := r.FormValue("path")
	errs := ta.teacherServices.UploadResource(models.Resources{SubjectId: uint(subjectId), Title: title, Description: description, Link: path})
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	http.StatusText(http.StatusOK)
}

func (ta *TeacherApi) DeleteResource(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	errs := ta.teacherServices.DeleteResource(uint(id))
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	http.StatusText(http.StatusOK)
}

func (ta *TeacherApi) ReportGrade(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	subjectId, _ := strconv.Atoi(r.FormValue("subjectid"))
	studentId, _ := strconv.Atoi(r.FormValue("studentid"))
	assessment, _ := strconv.Atoi(r.FormValue("assessment"))
	test, _ := strconv.Atoi(r.FormValue("test"))
	final, _ := strconv.Atoi(r.FormValue("final"))
	total, _ := strconv.Atoi(r.FormValue("total"))
	errs := ta.teacherServices.ReportGrade(models.Result{SubjectId: uint(subjectId), Id: uint(id), StudentId: uint(studentId), Assessment: assessment, Test: test, Final: final, Total: total})
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	http.StatusText(http.StatusOK)
}

func (ta *TeacherApi) ViewStudents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	students, errs := ta.teacherServices.ViewStudents(uint(id))
	output, _ := json.MarshalIndent(students, "", "\t\t")
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
