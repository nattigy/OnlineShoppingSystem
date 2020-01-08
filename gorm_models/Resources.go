package gorm_models

import "github.com/jinzhu/gorm"

type Resources struct {
	gorm.Model
	Id          int     `gorm:"primary_key"`
	SubjectId   Subject `gorm:"foreignkey:subject_id;association_foreignkey:id"`
	Title       string
	Description string
	path        string
}
