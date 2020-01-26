package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func Config() (*gorm.DB, error) {
	dbConn, err := gorm.Open("postgres", "host=localhost user=postgres dbname=pscsgorm password=mo port=2349 sslmode=disable")
	return dbConn, err
}
