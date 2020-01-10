package models

import "github.com/jinzhu/gorm"

type Resources struct {
	gorm.Model
	SubjectId   uint
	Subject     Subject
	Title       string
	Description string
	Path        string
}
