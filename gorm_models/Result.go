package gorm_models

type Result struct {
	Id         int     `gorm:"primary_key"`
	SubjectId  Subject `gorm:"foreignkey:subject_id;association_foreignkey:id"`
	StudentId  Student `gorm:"foreignkey:student_id;association_foreignkey:id"`
	Assessment int
	Test       int
	Final      int
	Total      int
}
