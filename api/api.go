package api

import (
	"encoding/json"
	"finbuild/database"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Response struct {
	Status  string                   `json:"status"` // success, fail, error
	Data    []map[string]interface{} `json:"data"`
	Message string                   `json:"message"`
}

type InfoAccounts struct {
	ID      int                      `json:"id"`
	Finance []map[string]interface{} `json:"finance"`
}

type Info struct {
	InfoUser InfoAccounts `json:"infoUser"`
}

func GetRoot(w http.ResponseWriter, r *http.Request) {

	// get balance
	type Account struct {
		Balance float64
	}
	var account Account
	// read database return and parse to struct
	q := r.URL.Query()
	id := q["id"]
	database.DB.Table("accounts").First(&account, "user_id = ?", id).Scan(&account).Find(&Account{})

	var msg []map[string]interface{}

	f := Account{Balance: account.Balance}

	type Users struct {
		ID string `json:"id"`
		FirstName string `json:"first_name"`
	}
	u := Users{}
	database.DB.Table("users").First(&u, "id = ?", id).Scan(&u).Find(&Users{})

	mp1 := map[string]interface{}{
		"user": u,
	}
	msg = append(msg, mp1)

	mp2 := map[string]interface{}{
		"finance": f,
	}
	msg = append(msg, mp2)

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
	w.Write(b)
}

// StartAPI Listen API with mux package
func StartAPI() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetRoot).Methods("GET")
	router.HandleFunc("/user", getUser).Methods("GET")
	router.HandleFunc("/exchange", postExchange).Methods("POST")

	fmt.Println("FinBuild API is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
