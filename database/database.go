package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=finbuild port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("ERROR: problem opening database", err)
	}
	DB = database
}
