package teacherApi

import (
	"encoding/json"
	"fmt"
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
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	er := ta.teacherServices.CreateTask(task)
	if er != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
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
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	newTask, er := ta.teacherServices.UpdateTask(task)
	if er != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	response, err := json.Marshal(newTask)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
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
	w.Header().Set("Content-Type", "application/json")
	var resource models.Resources
	err := json.NewDecoder(r.Body).Decode(&resource)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	er := ta.teacherServices.UploadResource(resource)
	if er != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
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
	w.Header().Set("Content-Type", "application/json")
	var result models.Result
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	er := ta.teacherServices.ReportGrade(result)
	if er != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
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
