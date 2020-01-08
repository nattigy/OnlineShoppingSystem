package gorm_models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Id               int
	Title            string
	Description      string
	ShortDescription string
	ClassRoomId      ClassRoom `gorm:"foreignkey:classroom_id;association_foreignkey:id"`
	ResourceId       Resources `gorm:"foreignkey:resource_id;association_foreignkey:id"`
	SubjectId        Subject   `gorm:"foreignkey:subject_id;association_foreignkey:id"`
	Deadline         string
}
