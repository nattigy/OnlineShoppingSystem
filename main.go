package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nattigy/parentschoolcommunicationsystem/delivery/http/api/admin"
	studentApi2 "github.com/nattigy/parentschoolcommunicationsystem/delivery/http/api/studentApi"
	teacherApi2 "github.com/nattigy/parentschoolcommunicationsystem/delivery/http/api/teacherApi"
	"github.com/nattigy/parentschoolcommunicationsystem/delivery/http/handlers/chatHandler"
	"github.com/nattigy/parentschoolcommunicationsystem/delivery/http/handlers/parentHandlers"
	"github.com/nattigy/parentschoolcommunicationsystem/delivery/http/handlers/studentHandlers"
	"github.com/nattigy/parentschoolcommunicationsystem/delivery/http/handlers/teacherHandlers"
	repository6 "github.com/nattigy/parentschoolcommunicationsystem/services/chatServices/repository"
	usecase6 "github.com/nattigy/parentschoolcommunicationsystem/services/chatServices/usecase"
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
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
	gormdb.CreateTable(&models.Admin{}, &models.Section{}, &models.Teacher{}, &models.Parent{}, &models.ClassRoom{}, &models.Subject{}, &models.Resources{}, &models.Student{}, &models.Result{}, &models.Task{}, &models.Comment{}, &models.User{}, &models.Session{}, &models.Message{})
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

	//Only run once
	//CreateTables(gormdb)
	//PopulateTables(gormdb)

	sessionRepo := repository5.NewSessionRepository(gormdb)
	sessionSer := usecase5.NewSessionUsecase(sessionRepo)

	studentRepo := repository.NewGormStudentRepository(gormdb)
	studentSer := usecase.NewStudentUsecase(studentRepo)
	studentHandler := studentHandlers.NewStudentHandler(templ, sessionSer, *studentSer)
	studentApi := studentApi2.NewStudentApi(studentSer)

	teacherRepo := repository2.NewGormTeacherRepository(gormdb)
	teacherSer := usecase2.NewTeacherUsecase(teacherRepo)
	teacherHandler := teacherHandlers.NewTeacherHandler(templ, sessionSer, *teacherSer)
	teacherApi := teacherApi2.NewTeacherApi(teacherSer)

	parentRepo := repository3.NewGormParentRepository(gormdb)
	parentSer := usecase3.NewParentUsecase(parentRepo)
	parentHandler := parentHandlers.NewParentHandler(templ, sessionSer, *parentSer)

	charRepo := repository6.NewChatRepository(gormdb)
	chatServ := usecase6.NewChatUsecase(charRepo)
	chatHandle := chatHandler.NewChatHandler(templ, chatServ, sessionSer, teacherSer, studentSer, parentSer)

	adminApi := admin.NewAdminApi(studentSer, teacherSer, parentSer)

	authRepo := repository4.NewAuthenticateRepository(gormdb)
	authSer := usecase4.NewAuthenticateUsecase(authRepo)
	authHandler := authenticationHandlers.NewAuthenticationHandler(templ, studentSer, teacherSer, parentSer, sessionSer, authSer)

	homeHandler := authenticationHandlers.NewHomePageHandler(templ, studentSer, teacherSer, parentSer, sessionSer)

	mux.HandleFunc("/", homeHandler.Home)
	mux.HandleFunc("/login", authHandler.Login)
	mux.Handle("/logout", authHandler.AuthenticateUser(http.HandlerFunc(authHandler.Logout)))

	mux.Handle("/student/viewTask", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.ViewTasks)))
	mux.Handle("/student/comment", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.Comment)))
	mux.Handle("/student/updateProfile", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.UpdateStudent)))
	mux.Handle("/student/viewClass", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.ViewClass)))
	mux.Handle("/student/resources", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.ViewResources)))
	mux.Handle("/student/viewResult", authHandler.AuthenticateUser(http.HandlerFunc(studentHandler.ViewResult)))

	mux.Handle("/teacher/makeNewPost", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.CreateTask)))
	mux.Handle("/teacher/editPost", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.UpdateTask)))
	mux.Handle("/teacher/removeTask", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.DeleteTask)))
	mux.Handle("/teacher/resource/delete", authHandler.AuthenticateUser(http.HandlerFunc(teacherHandler.DeleteResource)))
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

	//Chat
	mux.Handle("/parent/send", authHandler.AuthenticateUser(http.HandlerFunc(chatHandle.Send)))
	mux.Handle("/teacher/send", authHandler.AuthenticateUser(http.HandlerFunc(chatHandle.Send)))
	mux.Handle("/parent/receive", authHandler.AuthenticateUser(http.HandlerFunc(chatHandle.Get)))
	mux.Handle("/teacher/receive", authHandler.AuthenticateUser(http.HandlerFunc(chatHandle.Get)))

	//RestAPI

	router := httprouter.New()

	router.POST("/api/admin/student/new", adminApi.AddStudent)
	router.POST("/api/admin/teacher/new", adminApi.AddTeacher)
	router.POST("/api/admin/parent/new", adminApi.AddParent)
	router.GET("/api/admin/students", adminApi.GetStudents)
	router.GET("/api/admin/student/:id", adminApi.GetStudentById)
	router.GET("/api/admin/teachers", adminApi.GetTeachers)
	router.GET("/api/admin/teacher/:id", adminApi.GetTeacherById)
	router.GET("/api/admin/parents", adminApi.GetParents)
	router.GET("/api/admin/parent/:id", adminApi.GetParentById)
	router.DELETE("/api/admin/student/delete/:id", adminApi.DeleteStudent)
	router.DELETE("/api/admin/teacher/delete/:id", adminApi.DeleteTeacher)
	router.DELETE("/api/admin/parent/delete/:id", adminApi.DeleteParent)

	router.GET("/api/student/tasks", studentApi.ViewTasks)
	router.POST("/api/student/comment/new", studentApi.Comment)
	router.GET("/api/student/classmates", studentApi.ViewClass)
	router.GET("/api/student/resources", studentApi.ViewResources)
	router.GET("/api/student/result/:studentId", studentApi.ViewResult)

	router.POST("/api/teacher/post/new", teacherApi.CreateTask)
	router.POST("/api/teacher/post/update/:id", teacherApi.UpdateTask)
	router.DELETE("/api/teacher/post/delete/:id", teacherApi.DeleteTask)
	router.POST("/api/teacher/resource/new", teacherApi.UploadResource)
	router.POST("/api/teacher/grade/new", teacherApi.ReportGrade)
	router.GET("/api/teacher/students", teacherApi.ViewStudents)
	router.GET("/api/teacher/posts", teacherApi.GetTasks)

	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("server error : ", err)
	}
}

func PopulateTables(gormdb *gorm.DB) {
	admin1 := models.Admin{Id: 80, FirstName: "Zeleke", MiddleName: "Akale", Email: "zele@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e"}
	classRoom := models.ClassRoom{Id: 60, GradeLevel: 12, HomeRoom: 40}
	classRoom2 := models.ClassRoom{Id: 61, GradeLevel: 10, HomeRoom: 41}
	subject := models.Subject{Id: 100, SubjectName: "Math", ClassRoomId: 60}
	subject2 := models.Subject{Id: 101, SubjectName: "Physics", ClassRoomId: 61}
	user3 := models.User{Id: 80, Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", Email: "zele@gmail.com", Role: "admin"}
	section1 := models.Section{Id: 300, Section: "A", ClassRoomId: 60}
	section2 := models.Section{Id: 301, Section: "B", ClassRoomId: 61}
	fmt.Println(gormdb.Create(&classRoom))
	fmt.Println(gormdb.Create(&classRoom2))
	fmt.Println(gormdb.Create(&subject))
	fmt.Println(gormdb.Create(&subject2))
	fmt.Println(gormdb.Create(&section1))
	fmt.Println(gormdb.Create(&section2))
	fmt.Println(gormdb.Create(&admin1))
	fmt.Println(gormdb.Create(&user3))
}
