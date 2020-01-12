package repository

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"testing"
)

func TestGormStudentRepository_ViewTasks(t *testing.T) {
	gormdb, _ := database.GormConfig()
	s := NewGormStudentRepository(gormdb)
	classRoom := models.ClassRoom{Id: 1}
	subject := models.Subject{Id: 1}
	tasks, err := s.ViewTasks(classRoom, subject)
	if len(err) != 0 {
		t.Fatal(err)
	}
	fmt.Println(tasks)
}

func TestGormStudentRepository_ViewClass(t *testing.T) {
	gormdb, _ := database.GormConfig()
	s := NewGormStudentRepository(gormdb)
	classRoom := models.ClassRoom{Id: 1}
	students, err := s.ViewClass(classRoom)
	if len(err) != 0 {
		t.Fatal(err)
	}
	fmt.Println(students)
}

func TestGormStudentRepository_Comment(t *testing.T) {
	gormdb, _ := database.GormConfig()
	s := NewGormStudentRepository(gormdb)
	task := models.Task{Id: 1}
	student := models.Student{Id: 1}
	err := s.Comment(task, student, "my new comment")
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestGormStudentRepository_StudentUpdateProfile(t *testing.T) {
	gormdb, _ := database.GormConfig()
	s := NewGormStudentRepository(gormdb)
	student := models.Student{Id: 1, Email: "myNewEmail@gmail.Com", Password: "9876", ProfilePic: "/upload/newImage.jpg"}
	err := s.StudentUpdateProfile(student)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestGormStudentRepository_ViewResources(t *testing.T) {
	gormdb, _ := database.GormConfig()
	s := NewGormStudentRepository(gormdb)
	subject := models.Subject{Id: 1}
	resources, err := s.ViewResources(subject)
	if len(err) != 0 {
		t.Fatal(err)
	}
	fmt.Println(resources)
}

func TestGormStudentRepository_ViewResult(t *testing.T) {
	gormdb, _ := database.GormConfig()
	s := NewGormStudentRepository(gormdb)
	student := models.Student{Id: 1}
	result, err := s.ViewResult(student)
	if len(err) != 0 {
		t.Fatal(err)
	}
	fmt.Println(result)
}
