package models

type Teacher struct {
	FirstName  string
	MiddleName string
	Subject    string
	HomeRoom   bool
	Email      string
	Password   string
}

func (t *Teacher) MakeNewPost(task Task) bool {
	return true
}

func (t *Teacher) EditPost(task Task) bool {
	return false
}

func (t *Teacher) RemoveTask(task Task) bool {
	return false
}

func (t *Teacher) UploadResource(subject Subject) {

}

func (t *Teacher) UpdateProfile() {

}

func (t *Teacher) ReportGrade(grade Grade) {

}

func (t *Teacher) ViewClasses() []Student {
	return nil
}
