package models

type Task struct {
	Id               int
	Title            string
	Description      string
	ShortDescription string
	ClassRoomId      int
	ResourceId       int
	SubjectId        int
	PostedDate       string
	Deadline         string
}
