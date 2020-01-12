package usecase

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacher/repository"
	"testing"
)

func TestTeacherUsecase_MakeNewPost(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := repository.NewGormTeacherRepository(gormdb)
	u := NewTeacherUsecase(m)
	task := models.Task{ClassRoomId: 1, ShortDescription: "do it", Description: "do your project based on your requirements", Title: "Project", Deadline: "december 20"}
	subject := models.Subject{Id: 1}
	err := u.MakeNewPost(task, subject)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestTeacherUsecase_EditPost(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := repository.NewGormTeacherRepository(gormdb)
	u := NewTeacherUsecase(m)
	task := models.Task{Id: 2, ClassRoomId: 1, ShortDescription: "do it", Description: "do your project based on your requirements", Title: "Project", Deadline: "december 20"}
	err := u.EditPost(task)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestTeacherUsecase_RemoveTask(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := repository.NewGormTeacherRepository(gormdb)
	u := NewTeacherUsecase(m)
	task := models.Task{Id: 1}
	err := u.RemoveTask(task)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestTeacherUsecase_ReportGrade(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := repository.NewGormTeacherRepository(gormdb)
	u := NewTeacherUsecase(m)
	result := models.Result{SubjectId: 1, Assessment: 20, Test: 20, Final: 50, Total: 90}
	student := models.Student{Id: 2}
	err := u.ReportGrade(result, student)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestTeacherUsecase_TeacherUpdateProfile(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := repository.NewGormTeacherRepository(gormdb)
	u := NewTeacherUsecase(m)
	teacher := models.Teacher{Id: 11, Email: "kebede@gmail.com", Password: "7894", ProfilePic: "/upload/profile/newImage.jpg"}
	err := u.TeacherUpdateProfile(teacher)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestTeacherUsecase_UploadResource(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := repository.NewGormTeacherRepository(gormdb)
	u := NewTeacherUsecase(m)
	result := models.Result{SubjectId: 1, Assessment: 20, Test: 20, Final: 50, Total: 90}
	student := models.Student{Id: 2}
	err := u.ReportGrade(result, student)
	if len(err) != 0 {
		t.Fatal(err)
	}
}

func TestTeacherUsecase_ViewClasses(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := repository.NewGormTeacherRepository(gormdb)
	u := NewTeacherUsecase(m)
	classRoom := models.ClassRoom{Id: 1}
	students, err := u.ViewClasses(classRoom)
	if len(err) != 0 {
		t.Fatal(err)
	}
	fmt.Println(students)
}

func TestTeacherUsecase_FetchPosts(t *testing.T) {
	gormdb, _ := database.GormConfig()
	m := repository.NewGormTeacherRepository(gormdb)
	u := NewTeacherUsecase(m)
	subject := models.Subject{Id: 1}
	tasks, err := u.FetchPosts(subject)
	if len(err) != 0 {
		t.Fatal(err)
	}
	fmt.Println(tasks)
}
