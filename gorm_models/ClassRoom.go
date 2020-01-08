package gorm_models

type ClassRoom struct {
	Id         int `gorm:"primary_key"`
	GradeLevel int
	Section    string
	HomeRoom   Teacher `gorm:"foreignkey:home_room;association_foreignkey:id"`
	Subjects   []Subject
}
