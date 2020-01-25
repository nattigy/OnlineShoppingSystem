package main

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/delivery/http/handlers/parentHandlers"
	"github.com/nattigy/parentschoolcommunicationsystem/delivery/http/handlers/studentHandlers"
	"github.com/nattigy/parentschoolcommunicationsystem/delivery/http/handlers/teacherHandlers"
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/delivery/http/handlers/authenticationHandlers"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices/repository"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices/usecase"

	repository3 "github.com/nattigy/parentschoolcommunicationsystem/services/parentServices/repository"
	usecase3 "github.com/nattigy/parentschoolcommunicationsystem/services/parentServices/usecase"
	repository5 "github.com/nattigy/parentschoolcommunicationsystem/services/session/repository"
	usecase5 "github.com/nattigy/parentschoolcommunicationsystem/services/session/usecase"
	repository2 "github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices/repository"
	usecase2 "github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices/usecase"
	repository4 "github.com/nattigy/parentschoolcommunicationsystem/services/utility/repository"
	usecase4 "github.com/nattigy/parentschoolcommunicationsystem/services/utility/usecase"
)

var templ = template.Must(template.ParseGlob("ui/templates/*.html"))

func CreateTables(gormdb *gorm.DB) {
	gormdb.CreateTable(&models.Teacher{}, &models.Parent{}, &models.ClassRoom{}, &models.Subject{}, &models.Resources{}, &models.Student{}, &models.Result{}, &models.Task{}, &models.Comment{}, &models.User{}, &models.Session{}, &models.Message{})
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	gormdb, err := database.Config()
	if err != nil {
		fmt.Println(err)
	}
	defer gormdb.Close()

	//CreateTables(gormdb)

	sessionRepo := repository5.NewSessionRepository(gormdb)
	sessionSer := usecase5.NewSessionUsecase(sessionRepo)

	studentRepo := repository.NewGormStudentRepository(gormdb)
	studentSer := usecase.NewStudentUsecase(studentRepo)
	studentHandler := studentHandlers.NewStudentHandler(templ, sessionSer)

	teacherRepo := repository2.NewGormTeacherRepository(gormdb)
	teacherSer := usecase2.NewTeacherUsecase(teacherRepo)
	teacherHandler := teacherHandlers.NewTeacherHandler()

	parentRepo := repository3.NewGormParentRepository(gormdb)
	parentSer := usecase3.NewParentUsecase(parentRepo)
	parentHandler := parentHandlers.NewParentHandler(templ, sessionSer)

	authRepo := repository4.NewAuthenticateRepository(gormdb)
	authSer := usecase4.NewAuthenticateUsecase(authRepo)
	authHandler := authenticationHandlers.NewAuthenticationHandler(templ, studentSer, teacherSer, parentSer, sessionSer, authSer)

	homeHandler := authenticationHandlers.NewHomePageHandler(templ, studentSer, teacherSer, parentSer, sessionSer)

	router := httprouter.New()

	router.GET("/", homeHandler.Home)
	router.POST("/login", authHandler.Login)
	router.GET("/logout", authHandler.Logout)

	//router.GET("/student/viewResult", authHandler.UserHandler(httprouter.Handle(studentHandler.ViewResult)))
	//router.GET("/student/viewTask", authHandler.UserHandler(httprouter.Handle(studentHandler.ViewTasks)))
	//router.POST("/student/comment", authHandler.UserHandler(httprouter.Handle(studentHandler.Comment)))
	//router.POST("/student/updateProfile", authHandler.UserHandler(httprouter.Handle(studentHandler.UpdateStudent)))
	//router.GET("/student/viewClass", authHandler.UserHandler(httprouter.Handle(studentHandler.ViewClass)))
	//router.GET("/student/resources", authHandler.UserHandler(httprouter.Handle(studentHandler.ViewResources)))
	//router.GET("/student/viewResult", authHandler.UserHandler(httprouter.Handle(studentHandler.ViewResult)))

	mux.Handle("/student/viewTask", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.ViewTasks)))
	mux.Handle("/student/comment", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.Comment)))
	mux.Handle("/student/updateProfile", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.UpdateStudent)))
	mux.Handle("/student/viewClass", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.ViewClass)))
	mux.Handle("/student/resources", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.ViewResources)))
	mux.Handle("/student/viewResult", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.ViewResult)))

	mux.Handle("/teacher/editPost", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.UpdateTask)))
	mux.Handle("/teacher/removeTask", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.DeleteTask)))
	mux.Handle("/teacher/uploadResources", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.UploadResource)))
	mux.Handle("/teacher/updateProfile", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.UpdateTeacher)))
	mux.Handle("/teacher/reportGrade", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.ReportGrade)))
	mux.Handle("/teacher/viewClasses", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.ViewStudents)))
	mux.Handle("/teacher/fetchPosts", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.GetTasks)))

	mux.Handle("/parent/viewGrade", authHandler.AuthenticateUser(http.HandlerFunc(parentHandler.ViewGrade)))

	mux.Handle("/admin/student/new", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.AddStudent)))
	mux.Handle("/admin/students", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.GetStudents)))
	mux.Handle("/admin/student/delete", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.DeleteStudent)))
	mux.Handle("/admin/teacher/new", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.AddTeacher)))
	mux.Handle("/admin/teachers", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.GetTeachers)))
	mux.Handle("/admin/teacher/delete", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.DeleteTeacher)))
	mux.Handle("/admin/parent/new", authHandler.AuthenticateUser(http.HandlerFunc(parentHandler.AddParent)))
	mux.Handle("/admin/parents", authHandler.AuthenticateUser(http.HandlerFunc(parentHandler.GetParents)))
	mux.Handle("/admin/parent/delete", authHandler.AuthenticateUser(http.HandlerFunc(parentHandler.DeleteParent)))

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("server error : ", err)
	}
}
