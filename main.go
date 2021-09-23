package main

import (
	"finbuild/api"
	"finbuild/database"
)

func main() {

	// Start database connection
	database.InitDatabase()

	// api.StartAPI run API Listen
	api.StartAPI()
}
