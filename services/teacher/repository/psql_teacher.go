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

func (t *PsqlTeacherRepository) MakeNewPost(task models.Task, c models.ClassRoom) error {
	return nil
}

func (t *PsqlTeacherRepository) EditPost(task models.Task) error {
	return nil
}

func (t *PsqlTeacherRepository) RemoveTask(task models.Task) error {
	return nil
}

func (t *PsqlTeacherRepository) UploadResource(resource models.Resources, s models.Subject, room models.ClassRoom) error {
	return nil
}

func (t *PsqlTeacherRepository) TeacherUpdateProfile(teacher models.Teacher) error {
	return nil
}

func (t *PsqlTeacherRepository) ReportGrade(grade models.Result, student models.Student) error {
	return nil
}

func (t *PsqlTeacherRepository) ViewClasses(room models.ClassRoom) ([]models.Student, error) {
	return nil, nil
}
