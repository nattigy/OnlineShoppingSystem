package database

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Config() *sql.DB {
	dbConn, _ := sql.Open("postgres", "postgres://postgres:1234@localhost/pscs?sslmode=disable")
	return dbConn
}

func GormConfig() (*gorm.DB, error) {
	dbConn, err := gorm.Open("postgres", "host=localhost user=postgres dbname=pscsgorm password=1234")
	return dbConn, err
}
