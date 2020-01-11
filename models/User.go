package models

type User struct {
	Id       uint `gorm:"primary_key;auto_increment:false"`
	Email    string
	Password string
	Role     string
	LoggedIn bool
}
