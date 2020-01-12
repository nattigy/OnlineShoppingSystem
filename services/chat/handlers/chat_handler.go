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

func NewChatHandler(templ *template.Template, chatUsecase usecases.ChatUsecase, session session.SessionUsecase, utility utility.Utility) *ChatHandler {
	return &ChatHandler{templ: templ, chatUsecase: chatUsecase, Session: session, utility: utility}
}

type ChatInfo struct {
	Message []models.Message
	User    models.User
	Parents []models.Parent
	Teacher models.Teacher
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
		if len(err) != 0 {
			fmt.Println(err)
			return
		}
		errs := c.chatUsecase.Store(models.Parent{Id: user.Id}, teacher, models.Message{MessageContent: data, From: "parent"})
		if len(err) != 0 {
			fmt.Println(errs)
			return
		}
		http.Redirect(w, r, "/parent/receive", http.StatusSeeOther)
	} else if user.Role == "teacher" {
		parentId := r.FormValue("parentId")
		id, _ := strconv.Atoi(parentId)
		errs := c.chatUsecase.Store(models.Parent{Id: uint(id)}, models.Teacher{Id: user.Id}, models.Message{MessageContent: data, From: "teacher"})
		if len(errs) != 0 {
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
	if user.Role == "parent" {
		teacher, err := c.utility.GetTeacherById(user.Id)
		if len(err) != 0 {
			fmt.Println(err)
		}
		messages, errs := c.chatUsecase.Get(models.Parent{Id: user.Id}, teacher)
		if len(err) != 0 {
			fmt.Println(errs)
		}
		in := ChatInfo{
			Message: messages,
			User:    user,
			Teacher: teacher,
		}
		_ = c.templ.ExecuteTemplate(w, "parentChatPage", in)
	} else if user.Role == "teacher" {
		parents, errs := c.utility.GetParentsByTeacherId(user.Id)
		if len(errs) != 0 {
			fmt.Println(errs)
		}
		parentId := r.FormValue("parentId")
		id, _ := strconv.Atoi(parentId)
		messages, errs := c.chatUsecase.Get(models.Parent{Id: uint(id)}, models.Teacher{Id: user.Id})
		in := ChatInfo{
			Message: messages,
			User:    user,
			Parents: parents,
		}
		_ = c.templ.ExecuteTemplate(w, "teacherChatPage", in)
	}
}
