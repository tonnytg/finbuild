package api

import (
	log "finbuild/pkg/logging"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Response it is a format to json response
type Response struct {
	Status  string                   `json:"status"` // success, fail, error
	Data    []map[string]interface{} `json:"data"`
	Message string                   `json:"message"`
}

// StartAPI Listen API with mux package
func StartAPI() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetRoot).Methods("GET")
	router.HandleFunc("/user", getUser).Methods("GET")
	router.HandleFunc("/user", registerUser).Methods("POST")
	router.HandleFunc("/exchange", postExchange).Methods("POST")

	fmt.Println("FinBuild API is working on port :8888")
	err := http.ListenAndServe(":8888", router)
	msg := fmt.Sprintf("error http server: %v", err)
	log.Msg("CRITICAL", msg)
}
