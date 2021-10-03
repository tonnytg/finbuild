package database

import (
	log "finbuild/pkg/logging"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase start database with QueryField allow get one or more fields
func InitDatabase() {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=finbuild port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{QueryFields: true})
	if err != nil {
		msg := fmt.Sprintf("error to connect database: %v", err)
		log.Msg("CRITICAL", msg)
	}
	DB = database
}