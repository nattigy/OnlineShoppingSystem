package gorm_models

type Student struct {
	Id          int
	FirstName   string
	MiddleName  string
	Email       string
	Password    string
	ClassRoomId int
	ClassRoom   ClassRoom
	ParentId    int
	Parent      Parent
	ProfilePic  string
	Result      []Result
}
