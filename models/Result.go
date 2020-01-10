package models

type Result struct {
	Id         uint
	SubjectId  uint
	Subject    Subject
	StudentId  uint
	Student    Student
	Assessment int
	Test       int
	Final      int
	Total      int
}
