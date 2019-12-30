package models

type Subject struct {
	Id          int
	Grade       int
	TeacherId   Teacher
	SubjectName string
}
