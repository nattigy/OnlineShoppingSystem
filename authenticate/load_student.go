package authenticate

import (
	"fmt"
	"html/template"
	"net/http"
)

type LoadStudent struct {
	templ *template.Template
}

func NewLoadStudent(t *template.Template) *LoadStudent {
	return &LoadStudent{
		templ: t,
	}
}

func (l *LoadStudent) Student(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookie("session"))
	err := l.templ.ExecuteTemplate(w, "studentPortal.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}
