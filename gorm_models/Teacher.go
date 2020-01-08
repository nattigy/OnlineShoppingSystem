package gorm_models

type Teacher struct {
	Id         int `gorm:"primary_key"`
	FirstName  string
	MiddleName string
	Email      string
	Password   string
	ProfilePic string
}
