package models

type Student struct {
	FirstName  string
	MiddleName string
	Email      string
	Class      string
	Password   string
	Result     Result
	ParentName Parent
	ProfilePic string
}

