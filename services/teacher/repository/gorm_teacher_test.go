package repository

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"testing"
)

func TestGormTeacherRepository_MakeNewPost(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := NewGormTeacherRepository(gormdb)
	task := models.Task{ClassRoomId: 1, ShortDescription: "do it", Description: "do your project based on your requirements", Title: "Project", Deadline: "december 20"}
	subject := models.Subject{Id: 1}
	err := m.MakeNewPost(task, subject)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestGormTeacherRepository_EditPost(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := NewGormTeacherRepository(gormdb)
	task := models.Task{Id: 2, ClassRoomId: 1, ShortDescription: "do it", Description: "do your project based on your requirements", Title: "Project", Deadline: "december 20"}
	err := m.EditPost(task)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestGormTeacherRepository_RemoveTask(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := NewGormTeacherRepository(gormdb)
	task := models.Task{Id: 1}
	err := m.RemoveTask(task)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestGormTeacherRepository_UploadResource(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := NewGormTeacherRepository(gormdb)
	resource := models.Resources{Title: "web reference book", Description: "read it care fully"}
	subject := models.Subject{Id: 1}
	classRoom := models.ClassRoom{Id: 1}
	err := m.UploadResource(resource, subject, classRoom)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestGormTeacherRepository_TeacherUpdateProfile(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := NewGormTeacherRepository(gormdb)
	teacher := models.Teacher{Id: 11, Email: "kebede@gmail.com", Password: "7894", ProfilePic: "/upload/profile/newImage.jpg"}
	err := m.TeacherUpdateProfile(teacher)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestGormTeacherRepository_ReportGrade(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := NewGormTeacherRepository(gormdb)
	result := models.Result{SubjectId: 1, Assessment: 20, Test: 20, Final: 50, Total: 90}
	student := models.Student{Id: 2}
	err := m.ReportGrade(result, student)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestGormTeacherRepository_ViewClasses(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := NewGormTeacherRepository(gormdb)
	classRoom := models.ClassRoom{Id: 1}
	students, err := m.ViewClasses(classRoom)
	if len(err) != 0 {
		t.Fatal(err)
	}
	fmt.Println(students)
}

func TestGormTeacherRepository_FetchPosts(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := NewGormTeacherRepository(gormdb)
	subject := models.Subject{Id: 1}
	tasks, err := m.FetchPosts(subject)
	if len(err) != 0 {
		t.Fatal(err)
	}
	fmt.Println(tasks)
}
