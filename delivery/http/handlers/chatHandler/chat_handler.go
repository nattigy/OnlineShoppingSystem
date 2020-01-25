package chatHandler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/chatServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/parentServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices"
	"html/template"
	"net/http"
	"strconv"
)

type ChatHandler struct {
	templ          *template.Template
	chatServices   chatServices.ChatUsecase
	Session        session.SessionUsecase
	teacherUsecase teacherServices.TeacherUsecase
	studentUsecase studentServices.StudentUsecase
	parentUsecase  parentServices.ParentUsecase
}

func NewChatHandler(templ *template.Template, chatServices chatServices.ChatUsecase, session session.SessionUsecase, teacherUsecase teacherServices.TeacherUsecase, studentUsecase studentServices.StudentUsecase, parentUsecase parentServices.ParentUsecase) *ChatHandler {
	return &ChatHandler{templ: templ, chatServices: chatServices, Session: session, teacherUsecase: teacherUsecase, studentUsecase: studentUsecase, parentUsecase: parentUsecase}
}

type ChatInfo struct {
	Message []models.Message
	User    models.User
	Parents []models.Student
	Teacher models.Teacher
}

func (c *ChatHandler) Send(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	data := r.FormValue("message")
	if user.Role == "parent" {
		s, err := c.parentUsecase.GetChild(user.Id)
		teacher, err := c.studentUsecase.GetHomeRoomTeacher(s.Id)
		if len(err) != 0 {
			fmt.Println(err)
			return
		}
		errs := c.chatServices.Store(user.Id, teacher.Id, data, "parent")
		if len(err) != 0 {
			fmt.Println(errs)
			return
		}
		http.Redirect(w, r, "/parent/receive", http.StatusSeeOther)
	} else if user.Role == "teacher" {
		parentId := r.FormValue("parentId")
		id, _ := strconv.Atoi(parentId)
		errs := c.chatServices.Store(uint(id), user.Id, data, "teacher")
		if len(errs) != 0 {
			fmt.Println(errs)
			return
		}
		http.Redirect(w, r, "/teacher/receive", http.StatusSeeOther)
	}
}

func (c *ChatHandler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	if user.Role == "parent" {
		s, err := c.parentUsecase.GetChild(user.Id)
		teacher, err := c.studentUsecase.GetHomeRoomTeacher(s.Id)
		if len(err) != 0 {
			fmt.Println(err)
		}
		messages, errs := c.chatServices.Get(user.Id, teacher.Id)
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
		teacher, _ := c.teacherUsecase.GetTeacherById(user.Id)
		parents, errs := c.teacherUsecase.ViewStudents(teacher.ClassRoomId)
		if len(errs) != 0 {
			fmt.Println(errs)
		}
		parentId := r.FormValue("parentId")
		id, _ := strconv.Atoi(parentId)
		messages, errs := c.chatServices.Get(uint(id), user.Id)
		in := ChatInfo{
			Message: messages,
			User:    user,
			Parents: parents,
		}
		_ = c.templ.ExecuteTemplate(w, "teacherChatPage", in)
	}
}
