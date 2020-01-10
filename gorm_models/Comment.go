package gorm_models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Data      string
	TaskId    int
	Task      Task
	StudentId int
	Student   Student
}
