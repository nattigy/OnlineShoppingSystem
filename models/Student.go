package models

type Student struct {
	Id         int
	FirstName  string
	MiddleName string
	Email      string
	Password   string
	ClassRoom  ClassRoom
	ParentName Parent
	ProfilePic string
}
