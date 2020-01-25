package utility

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type SubjectUsecase interface {
	CreateSubject(newSubject models.Subject) []error
	GetSubjectById(id uint) (models.Subject, []error)
	DeleteSubject(id uint) []error
}

type ClassRoomUsecase interface {
	CreateClassRoom(newClassRoom models.ClassRoom) []error
	GetClassRoomById(id uint) (models.ClassRoom, []error)
	DeleteClassRoom(id uint) []error
}

type ResultUsecase interface {
	Update(newResult models.Result) (models.Result, []error)
	Delete(id uint) []error
}

type AuthenticationUsecase interface {
	Authenticate(u models.User) (bool, models.User, error)
}
