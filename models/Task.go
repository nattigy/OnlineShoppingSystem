package models

import "time"

type Task struct {
	Id               uint `gorm:"primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
	Title            string
	Description      string
	ShortDescription string
	SubjectId        uint
	Subject          Subject
	ClassRoomId      uint
	ClassRoom        ClassRoom
	Deadline         string
	Comments         []Comment
}
