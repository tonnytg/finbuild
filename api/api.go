package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StartAPI Listen API with mux package
func StartAPI() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetRoot).Methods("GET")
	router.HandleFunc("/user", getUser).Methods("GET")
	fmt.Println("FinBuild API is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func GetRoot(w http.ResponseWriter, r *http.Request) {

	type InfoAccounts struct {
		ID      int
		Finance []map[string]interface{}
	}

	type Info struct {
		Message  string
		InfoUser InfoAccounts
	}

	var result []map[string]interface{}

	mp1 := map[string]interface{}{
		"Assets": "1.000,00",
		"ETF":    "2.000,00",
	}

	result = append(result, mp1)

	// struct
	info := Info{
		Message: "success",
		InfoUser: InfoAccounts{
			ID:      1,
			Finance: result,
		},
	}

	b, err := json.Marshal(info)
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
