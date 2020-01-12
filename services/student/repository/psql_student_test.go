package repository

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"testing"
)

func TestViewTasks(t *testing.T) {
	v := NewPsqlStudentRepository(database.Config())
	classRoom := models.ClassRoom{
		Id:         12,
		GradeLevel: 12,
		Section:    "a",
	}
	subject := models.Subject{
		Id: 31,
	}
	data, err := v.ViewTasks(classRoom, subject)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

func TestComment(t *testing.T) {
	v := NewPsqlStudentRepository(database.Config())
	student := models.Student{
		Id: 11,
	}
	task := models.Task{
		Id: 51,
	}
	err := v.Comment(task, student, "i commented on this")
	fmt.Println(err)
}

func TestStudentUpdateProfile(t *testing.T) {
	p := NewPsqlStudentRepository(database.Config())
	s := models.Student{
		Id:         11,
		Email:      "mahdi@gmai.com",
		ProfilePic: "profile/mahdi",
		Password:   "4321",
	}
	err := p.StudentUpdateProfile(s)
	if err != nil {
		t.Error("Failed to update profile")
	}
}

func TestViewClass(t *testing.T) {
	c := NewPsqlStudentRepository(database.Config())
	class := models.ClassRoom{
		Id: 12,
	}
	data, err := c.ViewClass(class)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

func TestViewResources(t *testing.T) {

}

func TestViewResult(t *testing.T) {

}
