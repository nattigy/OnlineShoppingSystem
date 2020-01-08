package gorm_models

type Subject struct {
	Id          int `gorm:"primary_key"`
	Grade       int
	TeacherId   Teacher `gorm:"foreignkey:teacher_id;association_foreignkey:id"`
	SubjectName string
}
