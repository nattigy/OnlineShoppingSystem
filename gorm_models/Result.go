package gorm_models

type Result struct {
	Id         int
	SubjectId  int
	Subject    Subject
	StudentId  int
	Student    Student
	Assessment int
	Test       int
	Final      int
	Total      int
}
