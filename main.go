package main

import (
	"finbuild/pkg/db"
	"finbuild/web/api"
)

func main() {

	// Start db connection
	// this pool connect permit fast transactions
	db.InitDatabase()

	// api.StartAPI run API Listen
	api.StartAPI()
}
