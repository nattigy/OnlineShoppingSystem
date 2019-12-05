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

func MakeNewPost(task models.Task) (bool, error) {
	return true, nil
}

func EditPost(task models.Task) (bool, error) {
	return false, nil
}

func RemoveTask(task models.Task) (bool, error) {
	return false, nil
}

func UploadResource(subject models.Subject) error {
	return nil
}

func TeacherUpdateProfile(teacher models.Teacher) error {
	return nil
}

func ReportGrade(grade models.Grade) error {
	return nil
}

func ViewClasses() ([]models.Student, error) {
	return nil, nil
}
