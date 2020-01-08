package gorm_models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Id         int       `json:"id" gorm:"primary_key"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	ClassRoom  ClassRoom `json:"class_room" gorm:"foreignkey:classroom_id;association_foreignkey:id"`
	ParentName Parent    `json:"parent_name" gorm:"foreignkey:parent_id;association_foreignkey:id"`
	ProfilePic string    `json:"profile_pic"`
}
