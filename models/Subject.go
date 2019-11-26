package models

import "time"

type Subject struct {
	Id         int
	Assessment int
	Test       int
	Final      int
	Resources  Resource
}

