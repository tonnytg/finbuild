package webserver

import "net/http"

func CallRoutes() {
	// CallRoutes is the main function that handles all the routes
	http.HandleFunc("/", IndexHandler)
}
