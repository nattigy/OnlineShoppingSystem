package request

import (
	"html/template"
	"net/http"
)

type StudentRequest struct {
}

var templ = template.Must(template.ParseGlob("client/templates/*.html"))

func NewStudentRequest(h *http.ServeMux) {
	h.HandleFunc("/student/viewTask", ViewTasks)
	h.HandleFunc("/student/comment", Comment)
	h.HandleFunc("/student/updateProfile", StudentUpdateProfile)
	h.HandleFunc("/student/viewClass", ViewClass)
	h.HandleFunc("/student/viewResources", ViewResources)
	h.HandleFunc("/student/viewResult", ViewResult)
}

func ViewTasks(w http.ResponseWriter, r *http.Request) {

	//data, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//
	//}
	_ = templ.ExecuteTemplate(w, "index.html", r.Body)
}

func Comment(w http.ResponseWriter, r *http.Request) {

}

func StudentUpdateProfile(w http.ResponseWriter, r *http.Request) {

}

func ViewClass(w http.ResponseWriter, r *http.Request) {

}

func ViewResources(w http.ResponseWriter, r *http.Request) {

}

func ViewResult(w http.ResponseWriter, r *http.Request) {

}
