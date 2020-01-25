package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateRepository struct {
	conn *gorm.DB
}

func NewAuthenticateRepository(conn *gorm.DB) *AuthenticateRepository {
	return &AuthenticateRepository{conn: conn}
}

func (a *AuthenticateRepository) Authenticate(u models.User) (bool, models.User, error) {
	user := models.User{}
	errs := a.conn.Where("email = ?", u.Email).First(&user).GetErrors()
	auth := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if auth != nil {
		return false, user, auth
	} else if user.Id == 0 {
		return false, user, nil
	} else if len(errs) != 0 {
		return false, user, errs[0]
	}
	return true, user, nil
}
