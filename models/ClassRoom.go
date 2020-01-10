package models

type ClassRoom struct {
	Id         uint
	GradeLevel int
	Section    string
	HomeRoom   uint
	Subjects   []Subject
	Student    []Student
	Teacher    Teacher
}
