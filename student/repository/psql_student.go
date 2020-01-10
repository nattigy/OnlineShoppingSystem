package repository

import (
	"database/sql"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"time"
)

type PsqlStudentRepository struct {
	conn *sql.DB
}

func NewPsqlStudentRepository(Conn *sql.DB) *PsqlStudentRepository {
	return &PsqlStudentRepository{conn: Conn}
}

func (ps *PsqlStudentRepository) ViewTasks(c models.ClassRoom, s models.Subject) ([]models.Task, error) {
	data, err := ps.conn.Query("SELECT * FROM task WHERE class_room_id=$1 AND subject_id=$2", c.Id, s.Id)
	var tasks []models.Task
	if err != nil {
		return tasks, err
	}
	var task models.Task

	for data.Next() {
		if err := data.Scan(&task.Id, &task.Title, &task.Description, &task.ShortDescription, &task.ClassRoomId, &task.ResourceId,
			&task.SubjectId, &task.PostedDate, &task.Deadline,
		); err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (ps *PsqlStudentRepository) Comment(t models.Task, s models.Student, d string) error {
	_, err := ps.conn.Exec("INSERT INTO comment(data, student_id, task_id, date) VALUES ($1, $2, $3, $4)", d, s.Id, t.Id, time.Now())

	return err
}

func (ps *PsqlStudentRepository) StudentUpdateProfile(student models.Student) error {
	_, err := ps.conn.Exec("UPDATE student SET email=$1, password=$2, profile_pic=$3 WHERE id=$4", student.Email,
		student.Password, student.ProfilePic, student.Id)
	return err
}

func (ps *PsqlStudentRepository) ViewClass(classRoom models.ClassRoom) ([]models.Student, error) {
	data, err := ps.conn.Query("SELECT first_name, middle_name, email FROM student WHERE class_room=$1", classRoom.Id)
	var students []models.Student
	if err != nil {
		return students, err
	}

	var s models.Student

	for data.Next() {
		if err := data.Scan(&s.FirstName, &s.MiddleName, &s.Email); err != nil {
			return students, err
		}
		students = append(students, s)
	}
	return students, err
}

func (ps *PsqlStudentRepository) ViewResources(subject models.Subject) ([]models.Resource, error) {
	data, err := ps.conn.Query("SELECT * FROM resources WHERE subject_id=$1", subject.Id)
	var resources []models.Resource
	if err != nil {
		return resources, err
	}
	var resource models.Resource

	for data.Next() {
		if err := data.Scan(&resource.SubjectId, &resource.Title, &resource.Description, &resource.Path); err != nil {
			return resources, err
		}
		resources = append(resources, resource)
	}

	return resources, err
}

func (ps *PsqlStudentRepository) ViewResult(s models.Student) ([]models.Result, error) {
	data, err := ps.conn.Query("SELECT * FROM result WHERE student_id=$1", s.Id)
	var results []models.Result
	if err != nil {
		return results, err
	}
	var result models.Result

	for data.Next() {
		if err := data.Scan(&result.StudentId, &result.SubjectId, &result.Assessment, &result.Test, &result.Final, &result.Total); err != nil {
			return results, err
		}
		results = append(results, result)

	}
	return results, err

	return nil, nil
}
