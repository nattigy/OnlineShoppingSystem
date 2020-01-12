package usecase

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	repository2 "github.com/nattigy/parentschoolcommunicationsystem/services/student/repository"
	"testing"
)

func TestViewTasks(t *testing.T) {
	gormdb, _ := database.GormConfig()
	studentRepo := repository2.NewGormStudentRepository(gormdb)
	v := NewStudentUsecase(studentRepo)
	classRoom := models.ClassRoom{Id: 1}
	subject := models.Subject{Id: 1}
	tasks, err := v.ViewTasks(classRoom, subject)
	if len(err) != 0 {
		t.Fatal(err)
	}
	fmt.Println(tasks)
}

func TestComment(t *testing.T) {
	gormdb, _ := database.GormConfig()
	studentRepo := repository2.NewGormStudentRepository(gormdb)
	v := NewStudentUsecase(studentRepo)
	task := models.Task{Id: 1}
	student := models.Student{Id: 1}
	err := v.Comment(task, student, "my new comment")
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestStudentUpdateProfile(t *testing.T) {
	gormdb, _ := database.GormConfig()
	studentRepo := repository2.NewGormStudentRepository(gormdb)
	v := NewStudentUsecase(studentRepo)
	student := models.Student{Id: 1, Email: "myNewEmail@gmail.Com", Password: "9876", ProfilePic: "/upload/newImage.jpg"}
	err := v.StudentUpdateProfile(student)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestViewClass(t *testing.T) {
	gormdb, _ := database.GormConfig()
	studentRepo := repository2.NewGormStudentRepository(gormdb)
	v := NewStudentUsecase(studentRepo)
	classRoom := models.ClassRoom{Id: 1}
	students, err := v.ViewClass(classRoom)
	if len(err) != 0 {
		t.Fatal(err)
	}
	fmt.Println(students)
}

func TestViewResources(t *testing.T) {
	gormdb, _ := database.GormConfig()
	studentRepo := repository2.NewGormStudentRepository(gormdb)
	v := NewStudentUsecase(studentRepo)
	subject := models.Subject{Id: 1}
	resources, err := v.ViewResources(subject)
	if len(err) != 0 {
		t.Fatal(err)
	}
	fmt.Println(resources)
}

func TestViewResult(t *testing.T) {
	gormdb, _ := database.GormConfig()
	studentRepo := repository2.NewGormStudentRepository(gormdb)
	v := NewStudentUsecase(studentRepo)
	student := models.Student{Id: 1}
	result, err := v.ViewResult(student)
	if len(err) != 0 {
		t.Fatal(err)
	}
	fmt.Println(result)
}
