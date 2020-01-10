package authenticationHandlers

import (
	"github.com/nattigy/parentschoolcommunicationsystem/parent"
	"github.com/nattigy/parentschoolcommunicationsystem/session"
	"github.com/nattigy/parentschoolcommunicationsystem/student"
	"github.com/nattigy/parentschoolcommunicationsystem/teacher"
)

type LoginHandler struct {
	student student.StudentUsecase
	teacher teacher.TeacherUsecase
	parent  parent.ParentUsecase
	session session.SessionUsecase
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}
