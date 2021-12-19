package webserver

import (
	"log"
	"net/http"
	"os"
)

func Start() {

	log.Println("Starting webserver")

	// import middleware routes
	CallRoutes()
	port := "8080"
	// if PORT is exported use it
	if os.Getenv("WEBSERVER_PORT") != "" {
		port = os.Getenv("WEBSERVER_PORT")
	}
	// if PORT is not exported use 8080 as default
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
