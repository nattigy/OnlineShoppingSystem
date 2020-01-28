package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ClassRoom struct {
	Id         uint
	GradeLevel int
	HomeRoom   uint
	Subjects   []Subject
}

type Comment struct {
	gorm.Model
	FirstName string
	Data      string
	TaskId    uint
	StudentId uint
}

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

type Parent struct {
	Id         uint
	FirstName  string
	MiddleName string
	Email      string
	Password   string
	ProfilePic string
}

type Admin struct {
	Id         uint
	FirstName  string
	MiddleName string
	Email      string
	Password   string
	ProfilePic string
}

type Resources struct {
	gorm.Model
	SubjectId   uint
	Title       string
	Description string
	Link        string
}

type Result struct {
	Id         uint
	SubjectId  uint
	StudentId  uint
	Assessment int
	Test       int
	Final      int
	Total      int
}

type Section struct {
	Id          uint
	Section     string
	ClassRoomId uint
	Student     []Student
}

type Session struct {
	gorm.Model
	Name   string
	Uuid   string `json:"uuid" gorm:"unique;not null"`
	UserID uint   `json:"user_id" gorm:"not null"`
	Email  string `json:"email" gorm:"not null"`
	Role   string
}

type Student struct {
	Id          uint
	FirstName   string
	MiddleName  string
	Email       string
	Password    string
	SectionId   uint
	ClassRoomId uint
	ParentId    uint
	ProfilePic  string
	Result      []Result
}

type Subject struct {
	Id          uint
	ClassRoomId uint
	SubjectName string
	Tasks       []Task
	Resources   []Resources
}

type Task struct {
	Id               uint `gorm:"primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
	Title            string
	Description      string
	ShortDescription string
	SubjectId        uint
	ClassRoomId      uint
	Deadline         string
	Comments         []Comment
}

type Teacher struct {
	Id          uint
	FirstName   string
	MiddleName  string
	Email       string
	Password    string
	ProfilePic  string
	SubjectId   uint
	ClassRoomId uint
}

type User struct {
	Id       uint `gorm:"primary_key;auto_increment:false"`
	Email    string
	Password string
	Role     string
	LoggedIn bool
}
