package models

import "time"

type Message struct {
	Id             uint
	TeacherId      uint
	ParentId       uint
	MessageContent string
	Date           time.Time
}
