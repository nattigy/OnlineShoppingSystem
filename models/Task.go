package models

import (
	"time"
)

type Task struct {
	Id               int
	Title            string
	Description      string
	ShortDescription string
	PostedDate       time.Duration
	DeadLine         time.Duration
	Comments         []Comment
	OnSubject        Subject
}

