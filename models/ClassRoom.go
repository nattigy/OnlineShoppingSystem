package models

type ClassRoom struct {
	Id         int
	GradeLevel int
	Section    string
	Task       Task
	HomeRoom   Teacher
	Subjects   []Subject
	Students   []Student
}
