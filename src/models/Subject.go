package models

import "time"

type Subject struct {
	Id         int
	Assessment int
	Test       int
	Final      int
	Resources  Resource
}

type Resource struct {
	Title       string
	Date        time.Duration
	Description string
}
