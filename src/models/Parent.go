package models

type Parent struct {
	FirstName  string
	MiddleName string
	Email      string
	Password   string
	Child      Student
}

func (p *Parent) ViewGrade() Student {
	return Student{}
}
