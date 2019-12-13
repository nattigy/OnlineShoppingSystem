package repository

import (
	"database/sql"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type PsqlTeacherRepository struct {
	conn *sql.DB
}

func NewPsqlTeacherRepository(Conn *sql.DB) *PsqlTeacherRepository {
	return &PsqlTeacherRepository{conn: Conn}
}

func MakeNewPost(task models.Task, c models.ClassRoom) error {
	return nil
}

func EditPost(task models.Task) error {
	return nil
}

func RemoveTask(task models.Task) error {
	return nil
}

func UploadResource(resource models.Resource, s models.Subject, room models.ClassRoom) error {
	return nil
}

func TeacherUpdateProfile(teacher models.Teacher) error {
	return nil
}

func ReportGrade(grade models.Result, student models.Student) error {
	return nil
}

func ViewClasses(room models.ClassRoom) ([]models.Student, error) {
	return nil, nil
}
