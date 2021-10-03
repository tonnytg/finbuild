package api

import (
	"encoding/json"
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

// JParse Parse text to json
func JParse(w http.ResponseWriter, j ...map[string]interface{}) {

	// create a map for json template return
	var msg []map[string]interface{}
	for x, _ :=  range j {
		msg = append(msg, j[x])
	}

	jSend := Response{
		Status:  "success",
		Data:    msg,
		Message: "test",
	}

	b, err := json.Marshal(jSend)
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

// StartAPI Listen API with mux package
func StartAPI() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetRoot).Methods("GET")
	router.HandleFunc("/user", getUser).Methods("GET")
	router.HandleFunc("/user", registerUser).Methods("POST")

	router.HandleFunc("/wallet", getWallet).Methods("GET")

	router.HandleFunc("/exchange", postExchange).Methods("POST")

	fmt.Println("FinBuild API is working on port :8888")
	err := http.ListenAndServe(":8888", router)
	msg := fmt.Sprintf("error http server: %v", err)
	log.Msg("CRITICAL", msg)
}
