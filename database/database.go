package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// InitDatabase start database with QueryField allow get one or more fields
func InitDatabase() {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=finbuild port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{QueryFields: true})
	if err != nil {
		log.Println("ERROR: problem opening database", err)
	}
	DB = database
}
