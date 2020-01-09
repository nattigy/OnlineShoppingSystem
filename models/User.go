package gorm_models

type User struct {
	Id int `json:'id' gorm:"primary_key;auto_increment:false"`
	//FirstName  string
	//MiddleName string
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	LoggedIn bool `json:"logged_in"`
	//ProfilePic string
}
