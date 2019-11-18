package models

type Student struct {
	FirstName  string
	MiddleName string
	Email      string
	Class      string
	Password   string
	Grade      Grade
	ParentName Parent
}

func (s *Student) ViewTasks(subject Subject) []Task {
	return nil
}

func (s *Student) Comment(t Task) []Comment {
	return nil
}

func (s *Student) UpdateProfile() {

}

func (s *Student) ViewClass() []Student {
	return nil
}

func ViewResources(subject Subject) []Resource {
	return nil
}
