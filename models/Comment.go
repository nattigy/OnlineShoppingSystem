package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Data      string
	TaskId    uint
	Task      Task
	StudentId uint
	Student   Student
}
