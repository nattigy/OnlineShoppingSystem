package gorm_models

type Subject struct {
	Id          int
	TeacherId   int
	Teacher     Teacher
	ClassRoomId int
	ClassRoom   ClassRoom
	SubjectName string
	Tasks       []Task
	Resources   []Resources
}
