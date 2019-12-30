package authenticate

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"net/http"
)

type User struct {
	Email    string
	Password string
	Role     string
}

func HandelLogin(w http.ResponseWriter, r *http.Request) {
	key1 := "email"
	key2 := "password"

	user := User{
		Email:    r.FormValue(key1),
		Password: r.FormValue(key2),
	}

	auth, role, err := Authentcate(user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(auth, "   ", role)
}

func Authentcate(u User) (bool, string, error) {
	data, err := database.Config().Query("SELECT role FROM userinfo WHERE email=$1 AND password=$2", u.Email, u.Password)
	if err != nil {
		return false, "", err
	}
	var role string
	for data.Next() {
		if err := data.Scan(&role); err != nil {
			return false, role, err
		}
	}
	if role == "" {
		return false, "", nil
	}
	return true, role, nil
}
