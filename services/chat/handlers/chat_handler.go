package handlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/chat/usecases"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"github.com/nattigy/parentschoolcommunicationsystem/services/utility"
	"html/template"
	"net/http"
	"strconv"
)

type ChatHandler struct {
	templ       *template.Template
	chatUsecase usecases.ChatUsecase
	Session     session.SessionUsecase
	utility     utility.Utility
}

func NewChatHandler(templ *template.Template, chatUsecase usecases.ChatUsecase, utility utility.Utility) *ChatHandler {
	return &ChatHandler{templ: templ, chatUsecase: chatUsecase, utility: utility}
}

func (c *ChatHandler) Send(w http.ResponseWriter, r *http.Request) {
	user, err := c.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}
	data := r.FormValue("message")
	if user.Role == "parent" {
		teacher, err := c.utility.GetTeacherById(user.Id)
		if err != nil {
			fmt.Println(err)
			return
		}
		errs := c.chatUsecase.Store(models.Parent{Id: user.Id}, teacher, data)
		if errs != nil {
			fmt.Println(errs)
			return
		}
		http.Redirect(w, r, "/parent/receive", http.StatusSeeOther)
	} else if user.Role == "teacher" {
		parentId := r.FormValue("parentId")
		id, _ := strconv.Atoi(parentId)
		errs := c.chatUsecase.Store(models.Parent{Id: uint(id)}, models.Teacher{Id: user.Id}, data)
		if errs != nil {
			fmt.Println(errs)
			return
		}
		http.Redirect(w, r, "/teacher/receive", http.StatusSeeOther)
	}
}

func (c *ChatHandler) Get(w http.ResponseWriter, r *http.Request) {
	user, err := c.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}

}
