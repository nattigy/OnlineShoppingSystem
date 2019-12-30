package main

import (
	_ "github.com/lib/pq"
	"github.com/nattigy/parentschoolcommunicationsystem/authenticate"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"html/template"
	"net/http"

	_parHandlers "github.com/nattigy/parentschoolcommunicationsystem/delivery/http/parentHandlers"
	_stuHandlers "github.com/nattigy/parentschoolcommunicationsystem/delivery/http/studentHandlers"
	_techHandlers "github.com/nattigy/parentschoolcommunicationsystem/delivery/http/teacherHandlers"
	_parRepo "github.com/nattigy/parentschoolcommunicationsystem/parent/repository"
	_parUsecase "github.com/nattigy/parentschoolcommunicationsystem/parent/usecase"
	_stuRepo "github.com/nattigy/parentschoolcommunicationsystem/student/repository"
	_stuUsecase "github.com/nattigy/parentschoolcommunicationsystem/student/usecase"
	_techRepo "github.com/nattigy/parentschoolcommunicationsystem/teacher/repository"
	_techUsecase "github.com/nattigy/parentschoolcommunicationsystem/teacher/usecase"
)

var templ = template.Must(template.ParseGlob("ui/templates/*.html"))

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	db := database.Config()
	defer db.Close()

	studentRepo := _stuRepo.NewPsqlStudentRepository(db)
	studentUsecase := _stuUsecase.NewStudentUsecase(studentRepo)
	stuHandlers := _stuHandlers.NewStudentHandler(templ, *studentUsecase)

	mux.HandleFunc("/student/viewTask", stuHandlers.ViewTasks)
	mux.HandleFunc("/student/comment", stuHandlers.Comment)
	mux.HandleFunc("/student/updateProfile", stuHandlers.StudentUpdateProfile)
	mux.HandleFunc("/student/viewClass", stuHandlers.ViewClass)
	mux.HandleFunc("/student/resources", stuHandlers.ViewResources)
	mux.HandleFunc("/student/viewResult", stuHandlers.ViewResult)

	teacherRepo := _techRepo.NewPsqlTeacherRepository(db)
	teacherUsecase := _techUsecase.NewTeacherUsecase(teacherRepo)
	techHandlers := _techHandlers.NewTeacherHandler(templ, *teacherUsecase)

	mux.HandleFunc("/teacher/makeNewPost", techHandlers.MakeNewPost)
	mux.HandleFunc("/teacher/editPost", techHandlers.EditPost)
	mux.HandleFunc("/teacher/removeTask", techHandlers.RemoveTask)
	mux.HandleFunc("/teacher/uploadResources", techHandlers.UploadResource)
	mux.HandleFunc("/teacher/updateProfile", techHandlers.TeacherUpdateProfile)
	mux.HandleFunc("/teacher/reportGrade", techHandlers.ReportGrade)
	mux.HandleFunc("/teacher/viewClasses", techHandlers.ViewClasses)

	parentRepo := _parRepo.NewPsqlParentRepository(db)
	parentUsecase := _parUsecase.NewParentUsecase(parentRepo)
	parHandlers := _parHandlers.NewParentHandler(templ, *parentUsecase)

	mux.HandleFunc("/parent/viewGrade", parHandlers.ViewGrade)

	mux.HandleFunc("/login", authenticate.HandelLogin)
	mux.HandleFunc("/", Home)

	_ = http.ListenAndServe(":3000", mux)
}

func Home(w http.ResponseWriter, r *http.Request) {
	_ = templ.ExecuteTemplate(w, "index.html", "here")
}
