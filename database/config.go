package database

import "database/sql"

func Config() (*sql.DB, error) {
	dbconn, err := sql.Open("postgres", "postgres://postgres/1234@localhost/dbname")
	return dbconn, err
}
