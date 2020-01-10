package gorm_models

import "github.com/jinzhu/gorm"

type Resources struct {
	gorm.Model
	SubjectId   int
	Subject     Subject
	Title       string
	Description string
	path        string
}
