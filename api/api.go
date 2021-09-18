package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {

	if call["message"] == "success" {
		resp := call
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := call
		json.NewEncoder(w).Encode(resp)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {

	user := make(map[string]interface{})
	user["myName"] = "test"

	apiResponse(user, w)
}

func StartAPI() {
	router := mux.NewRouter()
	router.HandleFunc("/user", getUser).Methods("GET")
	fmt.Println("FinBuild API is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
