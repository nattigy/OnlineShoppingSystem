package models

type Subject struct {
	Id          uint
	TeacherId   uint
	Teacher     Teacher
	ClassRoomId uint
	ClassRoom   ClassRoom
	SubjectName string
	Tasks       []Task
	Resources   []Resources
}
