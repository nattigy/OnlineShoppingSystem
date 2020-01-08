package gorm_models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	Data      string
	Date      time.Duration
	StudentId Student `gorm:"foreignkey:student_id;association_foreignkey:id"`
	TaskId    Task    `gorm:"foreignkey:task_id;association_foreignkey:id"`
}
