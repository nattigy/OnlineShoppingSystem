package gorm_models

type Parent struct {
	Id         int `gorm:"primary_key"`
	FirstName  string
	MiddleName string
	Email      string
	Password   string
	ProfilePic string
	StudentId  Student `gorm:"foreignkey:student_id;association_foreignkey:id"`
}
