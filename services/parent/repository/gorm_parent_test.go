package repository

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"testing"
)

func TestGormParentRepository_ViewGrade(t *testing.T) {
	gormdb, _ := database.GormConfig()
	v := NewGormParentRepository(gormdb)
	student := models.Student{Id: 1}
	result, err := v.ViewGrade(student)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
