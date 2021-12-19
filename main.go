package main

import (
	"finbuild/pkg/db"
	"finbuild/web/api"
	"finbuild/web/webserver"
)

func main() {

	// Start db connection
	// this pool connect permit fast transactions
	db.InitDatabase()

	// api.StartAPI run API Listen
	go api.StartAPI()

	// webserver.StartWebServer run WebServer Listen
	webserver.Start()
}
