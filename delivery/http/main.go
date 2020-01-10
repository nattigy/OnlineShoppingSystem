package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/nattigy/parentschoolcommunicationsystem/authenticate"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	_parHandlers "github.com/nattigy/parentschoolcommunicationsystem/delivery/http/parentHandlers"
	_stuHandlers "github.com/nattigy/parentschoolcommunicationsystem/delivery/http/studentHandlers"
	_techHandlers "github.com/nattigy/parentschoolcommunicationsystem/delivery/http/teacherHandlers"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	_parRepo "github.com/nattigy/parentschoolcommunicationsystem/parent/repository"
	_parUsecase "github.com/nattigy/parentschoolcommunicationsystem/parent/usecase"
	_stuRepo "github.com/nattigy/parentschoolcommunicationsystem/student/repository"
	_stuUsecase "github.com/nattigy/parentschoolcommunicationsystem/student/usecase"
	_techRepo "github.com/nattigy/parentschoolcommunicationsystem/teacher/repository"
	_techUsecase "github.com/nattigy/parentschoolcommunicationsystem/teacher/usecase"
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseGlob("ui/templates/*.html"))

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	db := database.Config()
	defer db.Close()

	gormdb, err := database.GormConfig()
	if err != nil {
		fmt.Println(err)
	}
	defer gormdb.Close()

	CreateTables(gormdb)

	//var task []models.Task
	//gormdb.Where("class_room_id = ? AND subject_id = ?", 1, 1).Find(&task)
	//fmt.Println(task)

	studentRepo := _stuRepo.NewPsqlStudentRepository(db)
	studentUsecase := _stuUsecase.NewStudentUsecase(studentRepo)
	stuHandlers := _stuHandlers.NewStudentHandler(templ, *studentUsecase)
	loadStu := authenticate.NewLoadStudent(templ)

	mux.HandleFunc("/student", loadStu.Student)
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

func CreateTables(gormdb *gorm.DB) {
	gormdb.CreateTable(&models.Teacher{})
	gormdb.CreateTable(&models.Parent{})
	gormdb.CreateTable(&models.ClassRoom{})
	gormdb.CreateTable(&models.Subject{})
	gormdb.CreateTable(&models.Resources{})
	gormdb.CreateTable(&models.Student{})
	gormdb.CreateTable(&models.Result{})
	gormdb.CreateTable(&models.Task{})
	gormdb.CreateTable(&models.Comment{})
	gormdb.CreateTable(&models.User{})

	fmt.Println(gormdb.Model(&models.Subject{}).AddForeignKey("teacher_id", "teachers(id)", "RESTRICT", "RESTRICT"))
	fmt.Println(gormdb.Model(&models.Subject{}).AddForeignKey("class_room_id", "class_rooms(id)", "RESTRICT", "RESTRICT"))
	fmt.Println(gormdb.Model(&models.ClassRoom{}).AddForeignKey("home_room", "teachers(id)", "RESTRICT", "RESTRICT"))
	fmt.Println(gormdb.Model(&models.Comment{}).AddForeignKey("student_id", "students(id)", "RESTRICT", "RESTRICT"))
	fmt.Println(gormdb.Model(&models.Comment{}).AddForeignKey("task_id", "tasks(id)", "RESTRICT", "RESTRICT"))
	fmt.Println(gormdb.Model(&models.Resources{}).AddForeignKey("subject_id", "subjects(id)", "RESTRICT", "RESTRICT"))
	fmt.Println(gormdb.Model(&models.Result{}).AddForeignKey("subject_id", "subjects(id)", "RESTRICT", "RESTRICT"))
	fmt.Println(gormdb.Model(&models.Result{}).AddForeignKey("student_id", "students(id)", "RESTRICT", "RESTRICT"))
	fmt.Println(gormdb.Model(&models.Student{}).AddForeignKey("class_room_id", "class_rooms(id)", "RESTRICT", "RESTRICT"))
	fmt.Println(gormdb.Model(&models.Student{}).AddForeignKey("parent_id", "parents(id)", "RESTRICT", "RESTRICT"))
	fmt.Println(gormdb.Model(&models.Task{}).AddForeignKey("subject_id", "subjects(id)", "RESTRICT", "RESTRICT"))
	fmt.Println(gormdb.Model(&models.Task{}).AddForeignKey("class_room_id", "class_rooms(id)", "RESTRICT", "RESTRICT"))

	teacher := models.Teacher{FirstName: "Amanuel", MiddleName: "Tadele", Email: "aman@gmail.com", Password: "1234"}
	teacher2 := models.Teacher{FirstName: "Abebe", MiddleName: "Kebede", Email: "abebe@gmail.com", Password: "1234"}
	parent := models.Parent{FirstName: "Dinsa", MiddleName: "Lemi", Email: "dinsa@gmail.com", Password: "1234"}
	parent2 := models.Parent{FirstName: "Yewondwosen", MiddleName: "Akale", Email: "yewond@gmail.com", Password: "1234"}
	classRoom := models.ClassRoom{GradeLevel: 12, Section: "A", HomeRoom: 1}
	classRoom2 := models.ClassRoom{GradeLevel: 10, Section: "B", HomeRoom: 2}
	subject := models.Subject{TeacherId: 1, SubjectName: "Math", ClassRoomId: 1}
	subject2 := models.Subject{TeacherId: 2, SubjectName: "Physics", ClassRoomId: 2}
	task := models.Task{Title: "Home Work", Description: "Do it", ShortDescription: "Just Do it or i will kill you", SubjectId: 1, ClassRoomId: 1}
	task2 := models.Task{Title: "CLass Work", Description: "Do it", ShortDescription: "Just Do it or i will kill you", SubjectId: 1, ClassRoomId: 1}
	student := models.Student{FirstName: "Nathnael", MiddleName: "Yewondwosen", Email: "natnael@gmail.com", Password: "1234", ClassRoomId: 1, ParentId: 2}
	student2 := models.Student{FirstName: "Moti", MiddleName: "Dinsa", Email: "moti@gmail.com", Password: "1234", ClassRoomId: 1, ParentId: 1}
	comment := models.Comment{StudentId: 1, TaskId: 1, Data: "nati commenting"}
	comment2 := models.Comment{StudentId: 1, TaskId: 1, Data: "moti commenting"}
	user1 := models.User{Id: 1, Password: "1234", Email: "nati@gmail.com"}

	fmt.Println(gormdb.Create(&teacher))
	fmt.Println(gormdb.Create(&teacher2))
	fmt.Println(gormdb.Create(&parent))
	fmt.Println(gormdb.Create(&parent2))
	fmt.Println(gormdb.Create(&classRoom))
	fmt.Println(gormdb.Create(&classRoom2))
	fmt.Println(gormdb.Create(&subject))
	fmt.Println(gormdb.Create(&subject2))
	fmt.Println(gormdb.Create(&task))
	fmt.Println(gormdb.Create(&task2))
	fmt.Println(gormdb.Create(&student))
	fmt.Println(gormdb.Create(&student2))
	fmt.Println(gormdb.Create(&comment))
	fmt.Println(gormdb.Create(&comment2))
	fmt.Println(gormdb.Create(&user1))
}
