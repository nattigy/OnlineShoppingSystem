package usecase

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/gorm_models"
	"github.com/nattigy/parentschoolcommunicationsystem/student/repository"
	"testing"
)

func TestViewTasks(t *testing.T) {
	studentRepo := repository.NewPsqlStudentRepository(database.Config())
	v := NewStudentUsecase(studentRepo)
	c := gorm_models.ClassRoom{
		Id:         12,
		GradeLevel: 12,
		Section:    "a",
	}
	s := gorm_models.Subject{
		Id: 31,
	}
	data, err := v.ViewTasks(c, s)
	if err != nil {
		fmt.Println(data)
	}
	fmt.Println(data)
}

func TestComment(t *testing.T) {

}

func TestStudentUpdateProfile(t *testing.T) {

}

func TestViewClass(t *testing.T) {

}

func TestViewResources(t *testing.T) {

}

func TestViewResult(t *testing.T) {

}
