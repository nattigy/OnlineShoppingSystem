package models

type Student struct {
	Id          uint
	FirstName   string
	MiddleName  string
	Email       string
	Password    string
	ClassRoomId uint
	ClassRoom   ClassRoom
	ParentId    uint
	Parent      Parent
	ProfilePic  string
	Result      []Result
}
