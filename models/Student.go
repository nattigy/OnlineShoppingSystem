package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ClassRoom  int    `json:"class_room"`
	ParentName Parent `json:"parent_name"`
	ProfilePic string `json:"profile_pic"`
}
