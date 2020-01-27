package chatHandler

import (
	"fmt"
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
	Messages []models.Message
	User     models.User
	Parents  []models.Parent
	Teacher  models.Teacher
}

func (c *ChatHandler) Send(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	data := r.FormValue("message")
	if sess.Role == "parent" {
		s, err := c.parentUsecase.GetChild(sess.UserID)
		teacher, err := c.studentUsecase.GetHomeRoomTeacher(s.Id)
		if len(err) != 0 {
			fmt.Println(err)
			return
		}
		errs := c.chatServices.Store(sess.UserID, teacher.Id, data, "parent")
		if len(err) != 0 {
			fmt.Println(errs)
			return
		}
		http.Redirect(w, r, "/parent/receive", http.StatusSeeOther)
	} else if sess.Role == "teacher" {
		parentId := r.FormValue("parentId")
		id, _ := strconv.Atoi(parentId)
		fmt.Println("parentId : ", id)
		errs := c.chatServices.Store(uint(id), sess.UserID, data, "teacher")
		if len(errs) != 0 {
			fmt.Println(errs)
			return
		}
		http.Redirect(w, r, "/teacher/receive?parentId="+parentId, http.StatusSeeOther)
	}
}

func (c *ChatHandler) Get(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	user := models.User{Id: sess.ID, Email: sess.Email, Role: sess.Role, LoggedIn: true}
	if sess.Role == "parent" {
		s, err := c.parentUsecase.GetChild(sess.UserID)
		teacher, err := c.studentUsecase.GetHomeRoomTeacher(s.Id)
		if len(err) != 0 {
			fmt.Println(err)
		}
		messages, errs := c.chatServices.Get(sess.UserID, teacher.Id)
		if len(err) != 0 {
			fmt.Println(errs)
		}
		in := ChatInfo{
			Messages: messages,
			User:     user,
			Teacher:  teacher,
		}
		_ = c.templ.ExecuteTemplate(w, "parentChatPage.layout", in)
	} else if sess.Role == "teacher" {
		teacher, _ := c.teacherUsecase.GetTeacherById(sess.UserID)
		students, errs := c.teacherUsecase.ViewStudents(teacher.ClassRoomId)
		if len(errs) != 0 {
			fmt.Println(errs)
		}
		parentId := r.FormValue("parentId")
		id, _ := strconv.Atoi(parentId)
		messages, errs := c.chatServices.Get(uint(id), sess.UserID)
		var parents []models.Parent
		if len(students) > 0 {
			for i := 0; i < len(students); i++ {
				parent, _ := c.parentUsecase.GetParentById(students[i].ParentId)
				parents = append(parents, parent)
			}
		}
		if id == 0 {
			in := ChatInfo{
				Messages: []models.Message{},
				User:     user,
				Parents:  parents,
			}
			err := c.templ.ExecuteTemplate(w, "teacherChatPage.layout", in)
			if err != nil {
				fmt.Println(err)
			}
		}
		in := ChatInfo{
			Messages: messages,
			User:     user,
			Parents:  parents,
		}
		err := c.templ.ExecuteTemplate(w, "teacherChatPage.layout", in)
		if err != nil {
			fmt.Println(err)
		}
	}
}
