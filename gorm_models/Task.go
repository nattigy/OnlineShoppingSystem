package gorm_models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Title            string
	Description      string
	ShortDescription string
	SubjectId        int
	Subject          Subject
	ClassRoomId      int
	ClassRoom        ClassRoom
	Deadline         string
	Comments         []Comment
}
