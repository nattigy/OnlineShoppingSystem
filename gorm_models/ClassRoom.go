package gorm_models

type ClassRoom struct {
	Id         int
	GradeLevel int
	Section    string
	HomeRoom   int
	Subjects   []Subject
	Student    []Student
	Teacher    Teacher
}
