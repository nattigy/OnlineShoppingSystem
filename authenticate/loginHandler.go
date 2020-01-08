package authenticate

import (
	"encoding/json"
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"net/http"
)

type User struct {
	Email    string
	Password string
	Role     Role
	Id       int
	Loggedin bool
}

type Role struct {
	Student bool
	Teacher bool
	Parent  bool
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

	user := User{
		Email:    r.FormValue(key1),
		Password: r.FormValue(key2),
	}

	auth, role, err := Authenticate(user)
	if err != nil {
		fmt.Println("authentication error", err)
		return
	}

	var stu models.Student

	if auth {

		fmt.Println(r.Cookie("session"))
		if role == student {
			data, err := database.Config().Query("SELECT * FROM student WHERE email=$1", user.Email)
			if err != nil {
				fmt.Println(err)
				return
			}
			var p int
			for data.Next() {
				if err := data.Scan(&stu.FirstName, &stu.MiddleName, &stu.Id, &p, &stu.ProfilePic, &stu.ClassRoom, &stu.Email); err != nil {
					fmt.Println(err)
					return
				}
			}
			fmt.Println(r.Cookie("session"))
			cookie, err := r.Cookie("session")
			val, _ := json.Marshal(stu)
			fmt.Println(string(val))
			if err != nil {
				cookie = &http.Cookie{
					Name:     "session",
					Value:    string(val),
					HttpOnly: true,
					Path:     "/",
				}
				http.SetCookie(w, cookie)
			}
			//fmt.Println(cookie)
			//c := &http.Cookie{
			//	Name:     "session",
			//	Value:    "",
			//	Path:     "/",
			//	Expires: time.Unix(0, 0),
			//
			//	HttpOnly: true,
			//}
			//http.SetCookie(w, c)
			fmt.Println(r.Cookie("session"))
			http.Redirect(w, r, "/student/viewTask", http.StatusSeeOther)
		}
	}

}

func Authenticate(u User) (bool, string, error) {
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
