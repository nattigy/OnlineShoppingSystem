package models

type ClassRoom struct {
	GradeLevel string
	Section    string
	Task       Task
	HomeRoom   Teacher
	Subjects   []Subject
	Students   []Student
}
