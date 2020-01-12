package repository

import (
	"database/sql"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type PsqlParentRepository struct {
	conn *sql.DB
}

func NewPsqlParentRepository(Conn *sql.DB) *PsqlParentRepository {
	return &PsqlParentRepository{conn: Conn}
}

func (p *PsqlParentRepository) ViewGrade(student models.Student) ([]models.Result, error) {
	return []models.Result{}, nil
}
