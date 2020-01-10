package authenticate

import (
	"encoding/json"
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"net/http"
	"time"
)

type Role struct {
	Student bool
	Teacher bool
	Parent  bool
}

type User struct {
	Email    string
	Password string
	Role     Role
	Id       int
	Loggedin bool
}

const (
	student string = "student"
	teacher string = "teacher"
	parent  string = "parent"
)

type cookieValue struct {
	email string
	id    string
}

func (c *cookieValue) toString() string {
	return "{email: '" + c.email + "' id: '" + c.id + "' }"
}

func HandelLogin(w http.ResponseWriter, r *http.Request) {
	key1 := "email"
	key2 := "password"

	user := models.User{
		Email:    r.FormValue(key1),
		Password: r.FormValue(key2),
	}

	auth, role, err := Authenticate(user)
	if err != nil {
		res, _ := w.Write([]byte("0"))
		fmt.Println(res, err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if auth {
		cookie, err := r.Cookie("session")
		role.LoggedIn = true
		var value []byte
		value, _ = json.Marshal(role)

		if err == nil {

			c := &http.Cookie{
				Name:    "session",
				Value:   "",
				Path:    "/",
				Expires: time.Unix(0, 0),

				HttpOnly: true,
			}
			http.SetCookie(w, c)

			fmt.Println(string(value))
			cookie = &http.Cookie{
				Name:     "session",
				Value:    string(value),
				HttpOnly: true,
				Path:     "/",
			}
			http.SetCookie(w, cookie)
		} else {
			fmt.Println(string(value))
			cookie = &http.Cookie{
				Name:     "session",
				Value:    string(value),
				HttpOnly: true,
				Path:     "/",
			}
			http.SetCookie(w, cookie)
		}

		if role.Role == student {
			//data, err := database.Config().Query("SELECT * FROM student WHERE email=$1", user.Email)
			//if err != nil {
			//	fmt.Println(err)
			//	return
			//}
			//var p int
			//for data.Next() {
			//	if err := data.Scan(&stu.FirstName, &stu.MiddleName, &stu.Id, &p, &stu.ProfilePic, &stu.ClassRoom, &stu.Email); err != nil {
			//		fmt.Println(err)
			//		return
			//	}
			//}
			http.Redirect(w, r, "/student/viewTask", http.StatusSeeOther)
		} else if role.Role == teacher {
			http.Redirect(w, r, "/teacher", http.StatusSeeOther)
		} else if role.Role == parent {
			http.Redirect(w, r, "/parent", http.StatusSeeOther)
		}
	}

}

func Authenticate(user models.User) (bool, models.User, error) {
	gormdb, _ := database.GormConfig()
	gormdb.Where("email = ? AND password = ?", user.Email, user.Password).Find(&user)
	//data, err := database.Config().Query("SELECT role FROM userinfo WHERE email=$1 AND password=$2", user.Email, user.Password)
	//if err != nil {
	//	return false, "", err
	//}
	//var role string
	//for data.Next() {
	//	if err := data.Scan(&role); err != nil {
	//		return false, role, err
	//	}
	//}
	//if role == "" {
	//	return false, "", nil
	//}
	if user.Id == 0 {
		return false, models.User{}, nil
	}
	return true, user, nil
}
