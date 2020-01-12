package models

import "time"

type Message struct {
	Id             uint `gorm:"primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
	TeacherId      uint
	ParentId       uint
	MessageContent string
	From           string
}
