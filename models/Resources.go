package models

type Resource struct {
	SubjectId   Subject
	Title       string
	Description string
	Id          int
	path        string
}
