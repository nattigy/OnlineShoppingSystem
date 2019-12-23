package models

import "time"

type Comment struct {
	Data      string
	Date      time.Duration
	StudentId Student
	TaskId    Task
}
